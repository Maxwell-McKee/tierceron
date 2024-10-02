package cursorlib

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	kv "github.com/hashicorp/vault-plugin-secrets-kv"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/hashicorp/vault/sdk/plugin"
	"github.com/trimble-oss/tierceron/atrium/vestibulum/pluginutil"
	"github.com/trimble-oss/tierceron/atrium/vestibulum/trcdb/opts/prod"
	"github.com/trimble-oss/tierceron/atrium/vestibulum/trcshbase"
	"github.com/trimble-oss/tierceron/buildopts/coreopts"
	"github.com/trimble-oss/tierceron/buildopts/cursoropts"
	"github.com/trimble-oss/tierceron/buildopts/memonly"
	"github.com/trimble-oss/tierceron/buildopts/memprotectopts"
	"github.com/trimble-oss/tierceron/pkg/capauth"
	"github.com/trimble-oss/tierceron/pkg/core"
	eUtils "github.com/trimble-oss/tierceron/pkg/utils"
)

var logger *log.Logger

func InitLogger(l *log.Logger) {
	logger = l
}

func ParseCursorFields(e *logical.StorageEntry, tokenMap *map[string]interface{}, logger *log.Logger) error {
	logger.Println("ParseCursorFields")

	if e != nil {
		tokenData := map[string]interface{}{}
		decodeErr := e.DecodeJSON(&tokenData)
		if decodeErr != nil {
			return decodeErr
		}
		for cursor, cursorAttributes := range cursorFields {
			var tokenPtr *string
			var token string
			tokenNameKey := cursor

			if cursorAttributes.KeepSecret {
				if _, ptrOk := tokenData[cursor].(*string); ptrOk {
					tokenPtr = tokenData[cursor].(*string)
					tokenNameKey = cursor + "ptr"
				} else if _, strOk := tokenData[cursor].(string); strOk {
					token = tokenData[cursor].(string)
					tokenPtr = &token
					tokenNameKey = cursor + "ptr"
				}
			} else {
				if _, strOk := tokenData[cursor].(string); strOk {
					token = tokenData[cursor].(string)
					tokenPtr = &token
				}
			}
			if memonly.IsMemonly() {
				memprotectopts.MemProtect(nil, tokenPtr)
			}
			(*tokenMap)[tokenNameKey] = tokenPtr
		}
	}
	if len(*tokenMap) == 0 {
		return errors.New("No data")
	}

	logger.Println("ParseCursorFields complete")
	vaddrCheck := ""
	if v, vOk := (*tokenMap)["vaddress"].(string); vOk {
		vaddrCheck = v
	}

	return pluginutil.ValidateVaddr(vaddrCheck, logger)
}

var environments []string = []string{"dev"}
var environmentsProd []string = []string{"staging"}
var cursorFields map[string]cursoropts.CursorFieldAttributes
var KvInitialize func(context.Context, *logical.InitializationRequest) error
var curatorPluginConfig map[string]interface{}

var createUpdateFunc func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) = func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	pluginName := cursoropts.BuildOptions.GetPluginName()
	logger.Printf("%s CreateUpdate\n", pluginName)

	// Check that some fields are given
	if len(req.Data) == 0 {
		//ctx.Done()
		return logical.ErrorResponse("missing data fields"), nil
	}

	response, err2 := PersistCursorFieldsToVault(ctx, req.Path, &req.Storage, logger)
	if err2 != nil {
		return response, err2
	}

	logger.Printf("%s CreateUpdate complete\n", pluginName)

	return &logical.Response{
		Data: map[string]interface{}{
			"message": "Cursor updated",
		},
	}, nil
}

func PersistCursorFieldsToVault(ctx context.Context, key string, storage *logical.Storage, logger *log.Logger) (*logical.Response, error) {
	logger.Printf("PersistToVault\n")
	if key == "" {
		logger.Printf("PersistToVault missing path\n")
		return logical.ErrorResponse("missing path"), nil
	}

	tapMap := map[string]string{}
	for cursor, cursorAttributes := range cursorFields {
		if cursorAttributes.KeepSecret {
			if cursorPtr, ptrOk := curatorPluginConfig[fmt.Sprintf("%sptr", cursor)].(*string); ptrOk {
				tapMap[cursor] = *cursorPtr
			} else {
				logger.Printf("PersistToVault missed required coding for field %s\n", cursor)
			}
		} else {
			if _, strOk := curatorPluginConfig[cursor].(string); strOk {
				tapMap[cursor] = curatorPluginConfig[cursor].(string)
			} else {
				logger.Printf("PersistToVault missed required coding for field %s\n", cursor)
			}
		}
	}

	// JSON encode the data
	buf, err := json.Marshal(tapMap)
	if err != nil {
		//ctx.Done()
		logger.Printf("PersistToVault encode failure\n")
		return nil, fmt.Errorf("json encoding failed: %v", err)
	}

	// Write out a new key
	entry := &logical.StorageEntry{
		Key:   key,
		Value: buf,
	}
	if err := (*storage).Put(ctx, entry); err != nil {
		//ctx.Done()
		logger.Printf("PersistToVault write failure\n")
		return nil, fmt.Errorf("failed to write: %v", err)
	}
	logger.Printf("PersistToVault complete\n")

	return nil, nil
}

func GenerateSchema(fields map[string]cursoropts.CursorFieldAttributes) map[string]*framework.FieldSchema {
	schema := map[string]*framework.FieldSchema{}
	for key, value := range fields {
		schema[key] = &framework.FieldSchema{
			Type:        framework.TypeString,
			Description: value.Description,
		}
	}
	return schema
}

func GetCursorPluginOpts(pluginName string, tlsProviderFunc func() (*tls.Config, error)) *plugin.ServeOpts {
	return &plugin.ServeOpts{
		BackendFactoryFunc: func(ctx context.Context, config *logical.BackendConfig) (logical.Backend, error) {
			// Access backend configuration if needed
			fmt.Println("Backend configuration:", config)

			bkv, err := kv.Factory(ctx, config)
			KvInitialize = bkv.(*kv.PassthroughBackend).InitializeFunc

			bkv.(*kv.PassthroughBackend).InitializeFunc = func(ctx context.Context, req *logical.InitializationRequest) error {
				logger.Println("TrcCursorInitialize init begun.")
				if memonly.IsMemonly() {
					logger.Println("Unlocking everything.")
					memprotectopts.MemUnprotectAll(nil)
				}
				queuedEnvironments := environments
				if prod.IsProd() {
					queuedEnvironments = environmentsProd
				}

				trcshDriverConfig, err := trcshbase.TrcshInitConfig("dev", "", "", true, logger)
				eUtils.CheckError(&core.CoreConfig{
					ExitOnFailure: true,
					Log:           logger,
				}, err, true)

				cursorFields = cursoropts.BuildOptions.GetCursorFields()

				// Get common configs for deployer class of plugin.
				curatorPluginConfig = coreopts.BuildOptions.InitPluginConfig(curatorPluginConfig)

				var curatorEnv string = ""
				// Read in existing vault data from all existing environments on startup...
				for _, env := range queuedEnvironments {
					logger.Println("Processing env: " + env)
					curatorEnv = env
					tokenData, sgErr := req.Storage.Get(ctx, env)

					if sgErr != nil || tokenData == nil {
						if sgErr != nil {
							logger.Println("Missing configuration data for env: " + env + " error: " + sgErr.Error())
						} else {
							logger.Println("Missing configuration data for env: " + env)
						}
						// Get secrets from curator.
						logger.Printf("Field loading begun.\n")
						for cursorField, _ := range cursorFields {
							logger.Printf("Loading field: %s\n", cursorField)
							secretFieldValue, err := capauth.PenseQuery(trcshDriverConfig, cursoropts.BuildOptions.GetCapCuratorPath(), cursorField)
							if err != nil {
								logger.Printf("Failed to retrieve wanted key: %s\n", cursorField)
								continue
							}
							switch cursorField {
							case "vaddress", "caddress":
								curatorPluginConfig[cursorField] = *secretFieldValue
							default:
								curatorPluginConfig[fmt.Sprintf("%sptr", cursorField)] = secretFieldValue
							}
						}
						logger.Printf("Field loading complete.\n")
						PersistCursorFieldsToVault(ctx, env, &req.Storage, logger)
					} else {
						ptError := ParseCursorFields(tokenData, &curatorPluginConfig, logger)

						if ptError != nil {
							logger.Println("Bad configuration data for env: " + env + " error: " + ptError.Error())
						}
					}
				}
				logger.Println("Plugin config complete.")

				cursoropts.BuildOptions.TapInit()

				// Clean up tap
				e := os.Remove(fmt.Sprintf("%strcsnap.sock", cursoropts.BuildOptions.GetCapPath()))
				if e != nil {
					logger.Println("Unable to refresh socket.  Uneccessary.")
				}

				// Establish tap and feather.
				initErr := pluginutil.PluginTapFeatherInit(trcshDriverConfig, curatorPluginConfig)
				if initErr != nil {
					logger.Printf("Missing config for env: %s error: %s\n", curatorEnv, initErr.Error())
				}

				logger.Println("TrcCursorInitialize complete.")

				if KvInitialize != nil {
					logger.Println("Entering KvInitialize...")
					return KvInitialize(ctx, req)
				}

				return nil
			}

			bkv.(*kv.PassthroughBackend).Paths = []*framework.Path{
				{
					Pattern:         "(dev|QA|staging|prod)",
					HelpSynopsis:    "Configure an access token.",
					HelpDescription: "Use this endpoint to configure the auth tokens required by trcvault.",

					Fields: GenerateSchema(cursorFields),
					Callbacks: map[logical.Operation]framework.OperationFunc{
						logical.ReadOperation:   bkv.(*kv.PassthroughBackend).Paths[0].Callbacks[logical.ReadOperation],
						logical.CreateOperation: createUpdateFunc,
						logical.UpdateOperation: createUpdateFunc,
					},
				},
			}

			if err != nil {
				logger.Print("%s had an error: %v", pluginName, err.Error())
			}

			return bkv, err
		},
		TLSProviderFunc: tlsProviderFunc,
		Logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Trace,
			Output:     logger.Writer(),
			JSONFormat: false,
		}),
	}
}
