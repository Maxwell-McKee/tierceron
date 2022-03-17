package xutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	vcutils "tierceron/trcconfig/utils"
	"tierceron/trcx/extract"
	"tierceron/utils"
	eUtils "tierceron/utils"
	"tierceron/vaulthelper/kv"

	"github.com/hashicorp/vault/api"
	"gopkg.in/yaml.v2"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup

var templateResultChan = make(chan *extract.TemplateResultData, 5)

// GenerateSeedsFromVaultRaw configures the templates in trc_templates and writes them to trcx
func GenerateSeedsFromVaultRaw(config *eUtils.DriverConfig, fromVault bool, templatePaths []string) (string, bool, string, error) {
	// Initialize global variables
	valueCombinedSection := map[string]map[string]map[string]string{}
	valueCombinedSection["values"] = map[string]map[string]string{}

	secretCombinedSection := map[string]map[string]map[string]string{}
	secretCombinedSection["super-secrets"] = map[string]map[string]string{}

	// Declare local variables
	templateCombinedSection := map[string]interface{}{}
	sliceTemplateSection := []interface{}{}
	sliceValueSection := []map[string]map[string]map[string]string{}
	sliceSecretSection := []map[string]map[string]map[string]string{}
	maxDepth := -1

	endPath := ""
	service := ""
	if len(config.IndexFilter) > 0 {
		service = config.IndexFilter[0]
	}
	multiService := false
	var mod *kv.Modifier

	filteredTemplatePaths := templatePaths[:0]
	if len(config.FileFilter) != 0 {
		for _, filter := range config.FileFilter {
			if !strings.HasSuffix(filter, ".tmpl") {
				filter = filter + ".tmpl"
			}
			for _, templatePath := range templatePaths {
				if strings.HasSuffix(templatePath, filter) {
					filteredTemplatePaths = append(filteredTemplatePaths, templatePath)
				}
			}
		}
	}
	if len(filteredTemplatePaths) > 0 {
		templatePaths = filteredTemplatePaths
		filteredTemplatePaths = filteredTemplatePaths[:0]
	}

	envVersion := strings.Split(config.Env, "_")
	if len(envVersion) != 2 {
		// Make it so.
		envVersion = eUtils.SplitEnv(config.Env)
	}
	env := envVersion[0]
	version := envVersion[1]

	if config.Token != "" && config.Token != "novault" {
		var err error
		mod, err = kv.NewModifier(config.Insecure, config.Token, config.VaultAddress, env, config.Regions, config.Log)
		if err != nil {
			eUtils.LogErrorObject(config, err, false)
		}
		mod.Env = env
		mod.Version = version
		if len(config.ProjectSections) > 0 {
			mod.ProjectIndex = config.ProjectSections
			mod.RawEnv = strings.Split(config.EnvRaw, "_")[0]
			mod.SectionName = config.SectionName
			mod.SubSectionValue = config.SubSectionValue
		}
	}

	if len(filteredTemplatePaths) > 0 {
		filteredTemplatePaths = utils.RemoveDuplicates(filteredTemplatePaths)
		templatePaths = filteredTemplatePaths
	}

	if config.GenAuth && mod != nil {
		_, err := mod.ReadData("apiLogins/meta")
		if err != nil {
			eUtils.LogInfo(config, "Cannot genAuth with provided token.")
			return "", false, "", eUtils.LogAndSafeExit(config, "", 1)
		}
	}

	if config.Token != "novault" && mod.Version != "0" { //If version isn't latest or is a flag
		var noCertPaths []string
		var certPaths []string
		for _, templatePath := range templatePaths { //Seperate cert vs normal paths
			if !strings.Contains(templatePath, "Common") {
				noCertPaths = append(noCertPaths, templatePath)
			} else {
				certPaths = append(certPaths, templatePath)
			}
		}

		if config.WantCerts { //Remove unneeded template paths
			templatePaths = certPaths
		} else {
			templatePaths = noCertPaths
		}

		project := ""
		if len(config.VersionFilter) > 0 {
			project = config.VersionFilter[0]
		}
		for _, templatePath := range templatePaths {
			_, service, _ := utils.GetProjectService(templatePath) //This checks for nested project names

			config.VersionFilter = append(config.VersionFilter, service) //Adds nested project name to filter otherwise it will be not found.
		}

		if config.WantCerts { //For cert version history
			config.VersionFilter = append(config.VersionFilter, "Common")
		}

		config.VersionFilter = utils.RemoveDuplicates(config.VersionFilter)
		mod.VersionFilter = config.VersionFilter
		versionMetadataMap := utils.GetProjectVersionInfo(config, mod)

		if versionMetadataMap == nil {
			return "", false, "", eUtils.LogAndSafeExit(config, fmt.Sprintf("No version data found - this filter was applied during search: %v\n", config.VersionFilter), 1)
		} else if version == "versionInfo" { //Version flag
			var masterKey string
			first := true
			for key := range versionMetadataMap {
				passed := false
				if config.WantCerts {
					for _, service := range mod.VersionFilter {
						if !passed && strings.Contains(key, "Common") && strings.Contains(key, service) && !strings.Contains(key, project) && !strings.HasSuffix(key, "Common") {
							if len(key) > 0 {
								keySplit := strings.Split(key, "/")
								config.VersionInfo(versionMetadataMap[key], false, keySplit[len(keySplit)-1], first)
								passed = true
								first = false
							}
						}
					}
				} else {
					if len(key) > 0 && len(masterKey) < 1 {
						masterKey = key
						config.VersionInfo(versionMetadataMap[masterKey], false, "", false)
						return "", false, "", eUtils.LogAndSafeExit(config, "Version info provided.", 1)
					}
				}
			}
			return "", false, "", eUtils.LogAndSafeExit(config, "Version info provided.", 1)
		} else { //Version bound check
			versionNumbers := utils.GetProjectVersions(config, versionMetadataMap)
			utils.BoundCheck(config, versionNumbers, version)
		}
	}

	//Reciever for configs
	go func(c *eUtils.DriverConfig) {
		for {
			select {
			case tResult := <-templateResultChan:
				if config.Env == tResult.Env && config.SubSectionValue == tResult.SubSectionValue {
					sliceTemplateSection = append(sliceTemplateSection, tResult.InterfaceTemplateSection)
					sliceValueSection = append(sliceValueSection, tResult.ValueSection)
					sliceSecretSection = append(sliceSecretSection, tResult.SecretSection)
					if tResult.TemplateDepth > maxDepth {
						maxDepth = tResult.TemplateDepth
						//templateCombinedSection = interfaceTemplateSection
					}
					wg.Done()
				} else {
					go func(tResult *extract.TemplateResultData) {
						templateResultChan <- tResult
					}(tResult)
				}
			default:
			}
		}
	}(config)

	commonPathFound := false
	for _, tPath := range templatePaths {
		if strings.Contains(tPath, "Common") {
			commonPathFound = true
		}
	}

	commonPaths := []string{}
	if config.Token != "" && commonPathFound {
		var commonMod *kv.Modifier
		var err error
		commonMod, err = kv.NewModifier(config.Insecure, config.Token, config.VaultAddress, config.Env, config.Regions, config.Log)
		if err != nil {
			eUtils.LogErrorObject(config, err, false)
		}
		envVersion := strings.Split(config.Env, "_")
		commonMod.Env = envVersion[0]
		commonMod.Version = envVersion[1]
		commonMod.Version = commonMod.Version + "***X-Mode"

		commonPaths, err = vcutils.GetPathsFromProject(config, commonMod, "Common")
		if err != nil {
			eUtils.LogErrorObject(config, err, false)
		}
		if len(commonPaths) > 0 && strings.Contains(commonPaths[len(commonPaths)-1], "!=!") {
			commonPaths = commonPaths[:len(commonPaths)-1]
		}
		commonMod.Close()
	}

	// Configure each template in directory
	if config.Token != "novault" {
		//
		// Checking for existence of values for service in vault.
		//
		if strings.Contains(config.EnvRaw, ".*") || len(config.ProjectSections) > 0 {
			anyServiceFound := false
			serviceFound := false
			var acceptedTemplatePaths []string
			for _, templatePath := range templatePaths {
				_, _, templatePath = vcutils.GetProjectService(templatePath)
				_, _, indexed, _ := kv.PreCheckEnvironment(mod.Env)
				//This checks whether a enterprise env has the relevant project otherwise env gets skipped when generating seed files.
				if (strings.Contains(mod.Env, ".") || len(config.ProjectSections) > 0) && !serviceFound {
					var listValues *api.Secret
					var err error
					if config.SectionKey == "/Index/" && len(config.ProjectSections) > 0 {
						listValues, err = mod.ListEnv("super-secrets/" + strings.Split(config.EnvRaw, ".")[0] + config.SectionKey + config.ProjectSections[0] + "/" + config.SectionName + "/" + config.SubSectionValue + "/")
					} else if len(config.ProjectSections) > 0 { //If eid -> look inside Index and grab all environments
						listValues, err = mod.ListEnv("super-secrets/" + strings.Split(config.EnvRaw, ".")[0] + config.SectionKey + config.ProjectSections[0] + "/" + config.SectionName)
					} else if indexed {
						listValues, err = mod.ListEnv("super-secrets/" + mod.Env + "/")
					} else {
						listValues, err = mod.ListEnv("values/" + mod.Env + "/") //Fix values to add to project to directory
					}
					if err != nil {
						eUtils.LogErrorObject(config, err, false)
					} else if listValues == nil {
						eUtils.LogInfo(config, "No values were returned under values/.")
					} else {
						serviceSlice := make([]string, 0)
						for _, valuesPath := range listValues.Data {
							for _, serviceInterface := range valuesPath.([]interface{}) {
								serviceFace := serviceInterface.(string)
								if version != "0" {
									versionMap := utils.GetProjectVersionInfo(config, mod) //("super-secrets/" + strings.Split(config.EnvRaw, ".")[0] + config.SectionKey + config.ProjectSections[0] + "/" + config.SectionName + "/" + config.SubSectionValue + "/" + serviceFace)
									versionNumbers := utils.GetProjectVersions(config, versionMap)
									utils.BoundCheck(config, versionNumbers, version)
								}
								serviceSlice = append(serviceSlice, serviceFace)
							}
						}
						for _, listedService := range serviceSlice {
							if strings.TrimSuffix(listedService, "/") == service {
								serviceFound = true
							}
						}
					}
				}
				if serviceFound { //Exit for irrelevant enterprises
					acceptedTemplatePaths = append(acceptedTemplatePaths, templatePath)
					anyServiceFound = true
					serviceFound = false
				}
			}

			if !anyServiceFound { //Exit for irrelevant enterprises
				return "", false, "", eUtils.LogAndSafeExit(config, "No relevant services were found for this environment: "+mod.Env, 1)
			} else {
				if len(acceptedTemplatePaths) > 0 {
					// template paths further trimmed by vault.
					templatePaths = acceptedTemplatePaths
				}
			}
		}
	}

	var iFilterTemplatePaths []string
	if len(config.IndexFilter) > 0 {
		for _, iFilter := range config.IndexFilter {
			for _, tPath := range templatePaths {
				if strings.Contains(tPath, iFilter) {
					iFilterTemplatePaths = append(iFilterTemplatePaths, tPath)
				}
			}
		}
		templatePaths = iFilterTemplatePaths
	}

	// Configure each template in directory
	for _, templatePath := range templatePaths {
		wg.Add(1)
		go func(tp string, multiService bool, c *eUtils.DriverConfig, cPaths []string) {
			var project, service, env, version, innerProject string
			var errSeed error
			project = ""
			service = ""
			env = ""
			version = ""
			innerProject = "Not Found"

			// Map Subsections
			var templateResult extract.TemplateResultData
			var cds *vcutils.ConfigDataStore
			var goMod *kv.Modifier

			templateResult.ValueSection = map[string]map[string]map[string]string{}
			templateResult.ValueSection["values"] = map[string]map[string]string{}

			templateResult.SecretSection = map[string]map[string]map[string]string{}
			templateResult.SecretSection["super-secrets"] = map[string]map[string]string{}
			envVersion := eUtils.SplitEnv(config.Env)
			env = envVersion[0]
			version = envVersion[1]
			//check for template_files directory here
			project, service, tp = vcutils.GetProjectService(tp)

			if c.Token != "" && c.Token != "novault" {
				var err error
				goMod, err = kv.NewModifier(c.Insecure, c.Token, c.VaultAddress, c.Env, c.Regions, config.Log)
				if err != nil {
					eUtils.LogErrorObject(config, err, false)
					return
				}
				goMod.Env = env
				goMod.Version = version
				goMod.ProjectIndex = config.ProjectSections
				if len(goMod.ProjectIndex) > 0 {
					goMod.RawEnv = strings.Split(config.EnvRaw, "_")[0]
					goMod.SectionKey = config.SectionKey
					goMod.SectionName = config.SectionName
					goMod.SubSectionValue = config.SubSectionValue
				}
				if c.GenAuth {
					_, err := mod.ReadData("apiLogins/meta")
					if err != nil {
						eUtils.LogAndSafeExit(config, "Cannot genAuth with provided token.", -1)
						return
					}
				}

				cds = new(vcutils.ConfigDataStore)
				goMod.Version = goMod.Version + "***X-Mode"
				if goMod.SectionName != "" && goMod.SubSectionValue != "" {
					if goMod.SectionKey == "/Index/" {
						goMod.SectionPath = "super-secrets" + goMod.SectionKey + project + "/" + goMod.SectionName + "/" + goMod.SubSectionValue + "/" + service + config.SubSectionName
					} else {
						goMod.SectionPath = "super-secrets" + goMod.SectionKey + project + "/" + goMod.SectionName + "/" + goMod.SubSectionValue
					}
				}
				cds.Init(config, goMod, c.SecretMode, true, project, cPaths, service)

				if len(goMod.VersionFilter) >= 1 && strings.Contains(goMod.VersionFilter[len(goMod.VersionFilter)-1], "!=!") {
					// TODO: should this be before cds.Init???
					innerProject = strings.Split(goMod.VersionFilter[len(goMod.VersionFilter)-1], "!=!")[1]
					goMod.VersionFilter = goMod.VersionFilter[:len(goMod.VersionFilter)-1]
					if innerProject != "Not Found" {
						project = innerProject
						service = project
					}
				}

			}

			_, _, _, templateResult.TemplateDepth, errSeed = extract.ToSeed(config, goMod,
				cds,
				tp,
				project,
				service,
				fromVault,
				&(templateResult.InterfaceTemplateSection),
				&(templateResult.ValueSection),
				&(templateResult.SecretSection),
			)
			if errSeed != nil {
				eUtils.LogAndSafeExit(config, errSeed.Error(), -1)
				return
			}
			templateResult.Env = env + "_" + version
			templateResult.SubSectionValue = config.SubSectionValue
			templateResultChan <- &templateResult
		}(templatePath, multiService, config, commonPaths)
	}
	wg.Wait()

	// Combine values of slice
	CombineSection(config, sliceTemplateSection, maxDepth, templateCombinedSection)
	CombineSection(config, sliceValueSection, -1, valueCombinedSection)
	CombineSection(config, sliceSecretSection, -1, secretCombinedSection)

	var authYaml []byte
	var errA error

	// Add special auth section.
	if config.GenAuth {
		if mod != nil {
			connInfo, err := mod.ReadData("apiLogins/meta")
			if err == nil {
				authSection := map[string]interface{}{}
				authSection["apiLogins"] = map[string]interface{}{}
				authSection["apiLogins"].(map[string]interface{})["meta"] = connInfo
				authYaml, errA = yaml.Marshal(authSection)
				if errA != nil {
					eUtils.LogErrorObject(config, errA, false)
				}
			} else {
				return "", false, "", eUtils.LogAndSafeExit(config, "Attempt to gen auth for reduced privilege token failed.  No permissions to gen auth.", 1)
			}
		} else {
			authConfigurations := map[string]interface{}{}
			authConfigurations["authEndpoint"] = "<Enter Secret Here>"
			authConfigurations["pass"] = "<Enter Secret Here>"
			authConfigurations["sessionDB"] = "<Enter Secret Here>"
			authConfigurations["user"] = "<Enter Secret Here>"
			authConfigurations["trcAPITokenSecret"] = "<Enter Secret Here>"

			authSection := map[string]interface{}{}
			authSection["apiLogins"] = map[string]interface{}{}
			authSection["apiLogins"].(map[string]interface{})["meta"] = authConfigurations
			authYaml, errA = yaml.Marshal(authSection)
			if errA != nil {
				eUtils.LogErrorObject(config, errA, false)
			}
		}
	}

	// Create seed file structure
	template, errT := yaml.Marshal(templateCombinedSection)
	value, errV := yaml.Marshal(valueCombinedSection)
	secret, errS := yaml.Marshal(secretCombinedSection)

	if errT != nil {
		eUtils.LogErrorObject(config, errT, false)
	}

	if errV != nil {
		eUtils.LogErrorObject(config, errV, false)
	}

	if errS != nil {
		eUtils.LogErrorObject(config, errS, false)
	}
	templateData := string(template)
	// Remove single quotes generated by Marshal
	templateData = strings.ReplaceAll(templateData, "'", "")
	seedData := templateData + "\n\n\n" + string(value) + "\n\n\n" + string(secret) + "\n\n\n" + string(authYaml)

	return endPath, multiService, seedData, nil
}

// GenerateSeedsFromVault configures the templates in trc_templates and writes them to trcx
func GenerateSeedsFromVault(ctx eUtils.ProcessContext, config *eUtils.DriverConfig) (interface{}, error) {
	if config.Clean { //Clean flag in trcx
		if strings.HasSuffix(config.Env, "_0") {
			envVersion := eUtils.SplitEnv(config.Env)
			config.Env = envVersion[0]
		}
		_, err1 := os.Stat(config.EndDir + config.Env)
		err := os.RemoveAll(config.EndDir + config.Env)

		if err != nil {
			eUtils.LogErrorObject(config, err, false)
			eUtils.LogAndSafeExit(config, "", 1)
		}

		if err1 == nil {
			eUtils.LogInfo(config, "Seed removed from"+config.EndDir+config.Env)
		}
		return nil, nil
	}

	// Get files from directory
	tempTemplatePaths := []string{}
	for _, startDir := range config.StartDir {
		//get files from directory
		tp := GetDirFiles(startDir)
		tempTemplatePaths = append(tempTemplatePaths, tp...)
	}

	if len(tempTemplatePaths) == 0 {
		eUtils.LogErrorMessage(config, "No files found in trc_templates", true)
	}

	//Duplicate path remover
	keys := make(map[string]bool)
	templatePaths := []string{}
	for _, path := range tempTemplatePaths {
		if _, value := keys[path]; !value {
			keys[path] = true
			templatePaths = append(templatePaths, path)
		}
	}
	var mod *kv.Modifier

	if config.Token != "novault" {
		var err error
		// TODO: Redo/deleted the indexedEnv work...
		// Get filtered using mod and templates.
		mod, err = kv.NewModifier(config.Insecure, config.Token, config.VaultAddress, config.Env, config.Regions, config.Log)
		if err != nil {
			eUtils.LogErrorObject(config, err, false)
			return nil, eUtils.LogAndSafeExit(config, "", 1)
		}
		mod.Env = config.Env
	}
	templatePathsAccepted, err := eUtils.GetAcceptedTemplatePaths(config, mod, templatePaths)
	if err != nil {
		eUtils.LogErrorObject(config, err, false)
		eUtils.LogAndSafeExit(config, "", 1)
	}
	templatePaths = templatePathsAccepted

	endPath, multiService, seedData, errGenerateSeeds := GenerateSeedsFromVaultRaw(config, false, templatePaths)
	if errGenerateSeeds != nil {
		return errGenerateSeeds, nil
	}

	if endPath == "" && !multiService && seedData == "" {
		return nil, nil
	}

	suffixRemoved := ""
	envVersion := eUtils.SplitEnv(config.Env)
	config.Env = envVersion[0]
	if envVersion[1] != "0" {
		suffixRemoved = "_" + envVersion[1]
	}

	envBasePath, pathPart, pathInclude, _ := kv.PreCheckEnvironment(config.Env)

	if suffixRemoved != "" {
		config.Env = config.Env + suffixRemoved
	}

	if multiService {
		if strings.HasPrefix(config.Env, "local") {
			endPath = config.EndDir + "local/local_seed.yml"
		} else {
			if pathInclude {
				endPath = config.EndDir + envBasePath + "/" + pathPart + "/" + config.Env + "_seed.yml"
			} else {
				endPath = config.EndDir + envBasePath + "/" + config.Env + "_seed.yml"
			}
		}
	} else {
		if pathInclude {
			endPath = config.EndDir + envBasePath + "/" + pathPart + "/" + config.Env + "_seed.yml"
		} else if len(config.ProjectSections) > 0 {
			envBasePath, _, _, _ := kv.PreCheckEnvironment(config.EnvRaw)
			sectionNamePath := "/"
			subSectionValuePath := ""
			if config.SectionKey == "/Index/" {
				sectionNamePath = "/" + config.SectionName + "/"
				subSectionValuePath = config.SubSectionValue
			} else if config.SectionKey == "/Restricted/" {
				sectionNamePath = "/" + config.SectionName + "/"
				subSectionValuePath = config.Env
			}

			endPath = config.EndDir + envBasePath + config.SectionKey + config.ProjectSections[0] + sectionNamePath + subSectionValuePath + config.SubSectionName + "_seed.yml"
		} else {
			endPath = config.EndDir + envBasePath + "/" + config.Env + "_seed.yml"
		}
	}
	//generate template or certificate
	if config.WantCerts {
		var certData map[int]string
		certLoaded := false

		for _, templatePath := range templatePaths {

			project, service, templatePath := vcutils.GetProjectService(templatePath)

			envVersion := eUtils.SplitEnv(config.Env)

			mod, err := kv.NewModifier(config.Insecure, config.Token, config.VaultAddress, config.Env, config.Regions, config.Log)
			if err != nil {
				eUtils.LogErrorObject(config, err, false)
			}
			mod.Env = envVersion[0]
			mod.Version = envVersion[1]

			var ctErr error
			_, certData, certLoaded, ctErr = vcutils.ConfigTemplate(config, mod, templatePath, config.SecretMode, project, service, config.WantCerts, false)
			if ctErr != nil {
				if !strings.Contains(ctErr.Error(), "Missing .certData") {
					eUtils.CheckError(config, ctErr, true)
				}
			}

			if len(certData) == 0 {
				if certLoaded {
					eUtils.LogInfo(config, "Could not load cert "+templatePath)
					continue
				} else {
					continue
				}
			}

			certPath := fmt.Sprintf("%s", certData[2])
			eUtils.LogInfo(config, "Writing certificate: "+certPath+".")

			if strings.Contains(certPath, "ENV") {
				if len(mod.Env) >= 5 && (mod.Env)[:5] == "local" {
					envParts := strings.SplitN(mod.Env, "/", 3)
					certPath = strings.Replace(certPath, "ENV", envParts[1], 1)
				} else {
					certPath = strings.Replace(certPath, "ENV", mod.Env, 1)
				}
			}

			certDestination := config.EndDir + "/" + certPath
			certDestination = strings.ReplaceAll(certDestination, "//", "/")
			writeToFile(config, certData[1], certDestination)
			eUtils.LogInfo(config, "certificate written to "+certDestination)
		}
		return nil, nil
	}

	if config.Diff {
		if !strings.Contains(config.Env, "_") {
			config.Env = config.Env + "_0"
		}
		config.Update(&seedData, config.Env+"||"+config.Env+"_seed.yml")
	} else {
		writeToFile(config, seedData, endPath)
		// Print that we're done
		if strings.Contains(config.Env, "_0") {
			config.Env = strings.Split(config.Env, "_")[0]
		}
		if strings.Contains(envBasePath, "_") {
			envBasePath = strings.Split(envBasePath, "_")[0]
		}

		eUtils.LogInfo(config, "Seed created and written to "+endPath)
	}

	return nil, nil
}

func writeToFile(config *eUtils.DriverConfig, data string, path string) {
	byteData := []byte(data)
	//Ensure directory has been created
	dirPath := filepath.Dir(path)
	err := os.MkdirAll(dirPath, os.ModePerm)
	utils.CheckError(config, err, true)
	//create new file
	newFile, err := os.Create(path)
	utils.CheckError(config, err, true)
	//write to file
	_, err = newFile.Write(byteData)
	utils.CheckError(config, err, true)
	err = newFile.Sync()
	utils.CheckError(config, err, true)
	newFile.Close()
}

func GetDirFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	filePaths := []string{}
	//endPaths := []string{}
	if err != nil {
		//this is a file
		return []string{dir}
	}
	for _, file := range files {
		//add this directory to path names
		filename := file.Name()
		if strings.HasSuffix(filename, ".DS_Store") {
			continue
		}
		extension := filepath.Ext(filename)
		filePath := dir + file.Name()
		if !strings.HasSuffix(dir, "/") {
			filePath = dir + "/" + file.Name()
		}
		if extension == "" {
			//if subfolder add /
			filePath += "/"
		}
		//recurse to next level
		newPaths := GetDirFiles(filePath)
		filePaths = append(filePaths, newPaths...)
	}
	return filePaths
}

// MergeMaps - merges 2 maps recursively.
func MergeMaps(x1, x2 interface{}) interface{} {
	switch x1 := x1.(type) {
	case map[string]interface{}:
		x2, ok := x2.(map[string]interface{})
		if !ok {
			return x1
		}
		for k, v2 := range x2 {
			if v1, ok := x1[k]; ok {
				x1[k] = MergeMaps(v1, v2)
			} else {
				x1[k] = v2
			}
		}
	case nil:
		x2, ok := x2.(map[string]interface{})
		if ok {
			return x2
		}
	}
	return x1
}

// Combines the values in a slice, creating a singular map from multiple
// Input:
//	- slice to combine
//	- template slice to combine
//	- depth of map (-1 for value/secret sections)
func CombineSection(config *eUtils.DriverConfig, sliceSectionInterface interface{}, maxDepth int, combinedSectionInterface interface{}) {
	_, okMap := sliceSectionInterface.([]map[string]map[string]map[string]string)

	// Value/secret slice section
	if maxDepth < 0 && okMap {
		sliceSection := sliceSectionInterface.([]map[string]map[string]map[string]string)
		combinedSectionImpl := combinedSectionInterface.(map[string]map[string]map[string]string)
		for _, v := range sliceSection {
			for k2, v2 := range v {
				for k3, v3 := range v2 {
					if _, ok := combinedSectionImpl[k2][k3]; !ok {
						combinedSectionImpl[k2][k3] = map[string]string{}
					}
					for k4, v4 := range v3 {
						combinedSectionImpl[k2][k3][k4] = v4
					}
				}
			}
		}

		combinedSectionInterface = combinedSectionImpl

		// template slice section
	} else {
		if maxDepth < 0 && !okMap {
			eUtils.LogInfo(config, fmt.Sprintf("Env failed to gen.  MaxDepth: %d, okMap: %t\n", maxDepth, okMap))
		}
		sliceSection := sliceSectionInterface.([]interface{})

		for _, v := range sliceSection {
			MergeMaps(combinedSectionInterface, v)
		}
	}
}