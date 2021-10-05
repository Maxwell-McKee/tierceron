package trcxbase

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"tierceron/trcx/xutil"
	eUtils "tierceron/utils"

	"github.com/google/go-cmp/cmp"
)

type ResultData struct {
	inData *string
	inPath string
}

var resultMap = make(map[string]*string)
var envSlice = make([]string, 0)
var resultChannel = make(chan *ResultData, 5)
var envLength int
var mutex = &sync.Mutex{}

func messenger(inData *string, inPath string) {
	var data ResultData
	data.inData = inData
	data.inPath = inPath
	resultChannel <- &data
}

func reciever() {
	for {
		select {
		case data := <-resultChannel:
			if data != nil && data.inData != nil && data.inPath != "" {
				mutex.Lock()
				resultMap[data.inPath] = data.inData
				mutex.Unlock()
			}
		default:
		}
	}
}

func diffHelper() {
	keys := []string{}

	//Arranges keys for ordered output
	for _, env := range envSlice {
		keys = append(keys, env+"||"+env+"_seed.yml")
	}

	Reset := "\033[0m"
	Red := "\033[31m"
	Green := "\033[32m"
	Yellow := "\033[0;33m"

	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
	}

	keyA := keys[0]
	keyB := keys[1]
	keySplitA := strings.Split(keyA, "||")
	keySplitB := strings.Split(keyB, "||")
	mutex.Lock()
	envFileKeyA := resultMap[keyA]
	envFileKeyB := resultMap[keyB]
	mutex.Unlock()

	//Seperator
	if runtime.GOOS == "windows" {
		fmt.Printf("\n======================================================================================")
	} else {
		fmt.Printf("\n\033[1;35m======================================================================================\033[0m")
	}
	switch envLength {
	case 4:
		keyC := keys[2]
		keyD := keys[3]
		keySplitC := strings.Split(keyC, "||")
		keySplitD := strings.Split(keyD, "||")
		mutex.Lock()
		envFileKeyC := resultMap[keyC]
		envFileKeyD := resultMap[keyD]
		mutex.Unlock()

		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitB[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyB, envFileKeyA))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitC[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyC, envFileKeyA))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitD[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyD, envFileKeyA))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitB[0] + Reset + Green + " +Env-" + keySplitC[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyC, envFileKeyB))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitB[0] + Reset + Green + " +Env-" + keySplitD[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyD, envFileKeyB))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitC[0] + Reset + Green + " +Env-" + keySplitD[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyD, envFileKeyC))
	case 3:
		keyC := keys[2]
		keySplitC := strings.Split(keyC, "||")
		mutex.Lock()
		envFileKeyC := resultMap[keyC]
		mutex.Unlock()

		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitB[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyB, envFileKeyA))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitC[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyC, envFileKeyA))
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitB[0] + Reset + Green + " +Env-" + keySplitC[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyC, envFileKeyB))
	default:
		fmt.Print("\n" + Yellow + " (" + Reset + Red + "-Env-" + keySplitA[0] + Reset + Green + " +Env-" + keySplitB[0] + Reset + Yellow + ")" + Reset + "\n")
		fmt.Println(eUtils.LineByLineDiff(envFileKeyB, envFileKeyA))
	}

	//Seperator
	if runtime.GOOS == "windows" {
		fmt.Println("======================================================================================")
	} else {
		fmt.Println("\033[1;35m======================================================================================\033[0m")
	}
	keys = keys[:0] //Cleans keys for next file
}

func removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func versionHelper(versionData map[string]interface{}, templateOrValues bool, valuePath string) {
	Reset := "\033[0m"
	Cyan := "\033[36m"
	Red := "\033[31m"
	if runtime.GOOS == "windows" {
		Reset = ""
		Cyan = ""
		Red = ""
	}
	if versionData == nil {
		fmt.Println("No version data found for this environment")
		os.Exit(1)
	}
	for _, versionMetadata := range versionData {
		for field, data := range versionMetadata.(map[string]interface{}) {
			if field == "destroyed" && !data.(bool) {
				goto printOutput
			}
		}
	}
	return
printOutput:
	fmt.Println(Cyan + "======================================================================================" + Reset)
	keys := make([]int, 0, len(versionData))
	for versionNumber, _ := range versionData {
		versionNo, _ := strconv.ParseInt(versionNumber, 10, 64)
		keys = append(keys, int(versionNo))
	}
	sort.Ints(keys)
	for _, key := range keys {
		versionNumber := key
		versionMetadata := versionData[fmt.Sprint(key)]
		fields := make([]string, 0)
		fieldData := make(map[string]interface{}, 0)
		for field, data := range versionMetadata.(map[string]interface{}) {
			fields = append(fields, field)
			fieldData[field] = data
		}
		sort.Strings(fields)
		fmt.Println("Version " + fmt.Sprint(versionNumber) + " Metadata:")
		for _, field := range fields {
			fmt.Printf(field + ": ")
			fmt.Println(fieldData[field])
		}
		if keys[len(keys)-1] != versionNumber {
			fmt.Println(Red + "-------------------------------------------------------------------------------" + Reset)
		}
	}
	fmt.Println(Cyan + "======================================================================================" + Reset)
}

// CommonMain This executable automates the creation of seed files from template file(s).
// New seed files are written (or overwrite current seed files) to the specified directory.
func CommonMain(envPtr *string, addrPtrIn *string) {
	// Executable input arguments(flags)
	addrPtr := flag.String("addr", "", "API endpoint for the vault")
	if addrPtrIn != nil && *addrPtrIn != "" {
		addrPtr = addrPtrIn
	}
	startDirPtr := flag.String("startDir", "trc_templates", "Pull templates from this directory")
	endDirPtr := flag.String("endDir", "./trc_seeds/", "Write generated seed files to this directory")
	logFilePtr := flag.String("log", "./var/log/trcx.log", "Output path for log file")
	helpPtr := flag.Bool("h", false, "Provide options for trcx")
	tokenPtr := flag.String("token", "", "Vault access token")
	secretMode := flag.Bool("secretMode", true, "Only override secret values in templates?")
	genAuth := flag.Bool("genAuth", false, "Generate auth section of seed data?")
	cleanPtr := flag.Bool("clean", false, "Cleans seed files locally")
	secretIDPtr := flag.String("secretID", "", "Secret app role ID")
	appRoleIDPtr := flag.String("appRoleID", "", "Public app role ID")
	tokenNamePtr := flag.String("tokenName", "", "Token name used by this trcx to access the vault")
	noVaultPtr := flag.Bool("novault", false, "Don't pull configuration data from vault.")
	pingPtr := flag.Bool("ping", false, "Ping vault.")
	insecurePtr := flag.Bool("insecure", false, "By default, every ssl connection is secure.  Allows to continue with server connections considered insecure.")
	diffPtr := flag.Bool("diff", false, "Diff files")
	versionPtr := flag.Bool("version", false, "Gets version metadata information")

	// Checks for proper flag input
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		s := args[i]
		if s[0] != '-' {
			fmt.Println("Wrong flag syntax: ", s)
			os.Exit(1)
		}
	}

	flag.Parse()

	Yellow := "\033[33m"
	Reset := "\033[0m"
	if runtime.GOOS == "windows" {
		Reset = ""
		Yellow = ""
	}

	//check for clean + env flag
	cleanPresent := false
	envPresent := false
	for _, arg := range args {
		if strings.Contains(arg, "clean") {
			cleanPresent = true
		}
		if strings.Contains(arg, "env") {
			envPresent = true
		}
	}

	if cleanPresent && !envPresent {
		fmt.Println("Environment must be defined with -env=env1,... for -clean usage")
		os.Exit(1)
	} else if *diffPtr && *versionPtr {
		fmt.Println("-version flag cannot be used with -diff flag")
		os.Exit(1)
	}

	if *versionPtr {
		if strings.Contains(*envPtr, ",") {
			fmt.Println(Yellow + "Invalid environment, please specify one environment." + Reset)
			os.Exit(1)
		}
		envVersion := strings.Split(*envPtr, "_")
		if len(envVersion) > 1 && envVersion[1] != "" && envVersion[1] != "0" {

			fmt.Println(Yellow + "Specified versioning not available, using " + envVersion[0] + " as environment" + Reset)
		}
		envSlice = append(envSlice, *envPtr+"_versionInfo")
		goto skipDiff
	}

	//Diff flag parsing check
	if *diffPtr {
		if strings.ContainsAny(*envPtr, ",") { //Multiple environments
			*envPtr = strings.ReplaceAll(*envPtr, "latest", "0")
			envSlice = strings.Split(*envPtr, ",")
			envLength = len(envSlice)
			if len(envSlice) > 4 {
				fmt.Println("Unsupported number of environments - Maximum: 4")
				os.Exit(1)
			}
			for i, env := range envSlice {
				if env == "local" {
					fmt.Println("Unsupported env: local not available with diff flag")
					os.Exit(1)
				}
				if !strings.Contains(env, "_") {
					envSlice[i] = env + "_0"
				}
			}
		} else {
			fmt.Println("Incorrect format for diff: -env=env1,env2,...")
			os.Exit(1)
		}
	} else {
		if strings.ContainsAny(*envPtr, ",") {
			fmt.Println("-diff flag is required for multiple environments - env: -env=env1,env2,...")
			os.Exit(1)
		}
		envSlice = append(envSlice, (*envPtr))
		envVersion := strings.Split(*envPtr, "_") //Break apart env+version for token
		*envPtr = envVersion[0]
		eUtils.AutoAuth(*insecurePtr, secretIDPtr, appRoleIDPtr, tokenPtr, tokenNamePtr, envPtr, addrPtr, *pingPtr)
		if len(envVersion) >= 2 { //Put back env+version together
			*envPtr = envVersion[0] + "_" + envVersion[1]
			if envVersion[1] == "" {
				fmt.Println("Must declare desired version number after '_' : -env=env1_ver1")
				os.Exit(1)
			}
		} else {
			*envPtr = envVersion[0] + "_0"
		}
	}

skipDiff:
	// Prints usage if no flags are specified
	if *helpPtr {
		flag.Usage()
		os.Exit(1)
	}
	if _, err := os.Stat(*startDirPtr); os.IsNotExist(err) {
		fmt.Println("Missing required start template folder: " + *startDirPtr)
		os.Exit(1)
	}
	if _, err := os.Stat(*endDirPtr); os.IsNotExist(err) {
		fmt.Println("Missing required start seed folder: " + *endDirPtr)
		os.Exit(1)
	}

	// If logging production directory does not exist and is selected log to local directory
	if _, err := os.Stat("./var/log/"); os.IsNotExist(err) && *logFilePtr == "./var/log/trcx.log" {
		*logFilePtr = "./trcx.log"
	}

	regions := []string{}

	if len(envSlice) == 1 && !*noVaultPtr {
		if *envPtr == "staging" || *envPtr == "prod" {
			secretIDPtr = nil
			appRoleIDPtr = nil
			regions = eUtils.GetSupportedProdRegions()
		}
		eUtils.AutoAuth(*insecurePtr, secretIDPtr, appRoleIDPtr, tokenPtr, tokenNamePtr, envPtr, addrPtr, *pingPtr)
	}

	if tokenPtr == nil || *tokenPtr == "" && !*noVaultPtr && len(envSlice) == 1 {
		fmt.Println("Missing required auth token.")
		os.Exit(1)
	}

	if len(*envPtr) >= 5 && (*envPtr)[:5] == "local" {
		var err error
		*envPtr, err = eUtils.LoginToLocal()
		fmt.Println(*envPtr)
		eUtils.CheckError(err, true)
	}

	// Initialize logging
	f, err := os.OpenFile(*logFilePtr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	eUtils.CheckError(err, true)
	logger := log.New(f, "[trcx]", log.LstdFlags)
	logger.Println("=============== Initializing Seed Generator ===============")

	logger.SetPrefix("[trcx]")
	logger.Printf("Looking for template(s) in directory: %s\n", *startDirPtr)

	var waitg sync.WaitGroup
	go reciever() //Channel reciever
	for _, env := range envSlice {
		envVersion := strings.Split(env, "_") //Break apart env+version for token
		*envPtr = envVersion[0]
		*tokenPtr = ""
		if secretIDPtr != nil && *secretIDPtr != "" && appRoleIDPtr != nil && *appRoleIDPtr != "" {
			*tokenPtr = ""
		}
		if !*noVaultPtr {
			eUtils.AutoAuth(*insecurePtr, secretIDPtr, appRoleIDPtr, tokenPtr, tokenNamePtr, envPtr, addrPtr, *pingPtr)
		} else {
			*tokenPtr = "novault"
		}

		if len(envVersion) >= 2 { //Put back env+version together
			*envPtr = envVersion[0] + "_" + envVersion[1]
		} else {
			*envPtr = envVersion[0] + "_0"
		}
		config := eUtils.DriverConfig{
			Insecure:       *insecurePtr,
			Token:          *tokenPtr,
			VaultAddress:   *addrPtr,
			Env:            *envPtr,
			Regions:        regions,
			SecretMode:     *secretMode,
			ServicesWanted: []string{},
			StartDir:       append([]string{}, *startDirPtr),
			EndDir:         *endDirPtr,
			WantCert:       false,
			GenAuth:        *genAuth,
			Log:            logger,
			Clean:          *cleanPtr,
			Diff:           *diffPtr,
			Update:         messenger,
			VersionInfo:    versionHelper,
		}
		waitg.Add(1)
		go func() {
			defer waitg.Done()
			eUtils.ConfigControl(config, xutil.GenerateSeedsFromVault)
		}()
	}
	waitg.Wait()
	close(resultChannel)
	if *diffPtr { //Diff if needed
		waitg.Add(1)
		go func() {
			defer waitg.Done()
			diffHelper()
		}()
	}
	waitg.Wait() //Wait for diff

	logger.SetPrefix("[trcx]")
	logger.Println("=============== Terminating Seed Generator ===============")
	logger.SetPrefix("[END]")
	logger.Println()

	// Terminate logging
	f.Close()
}
