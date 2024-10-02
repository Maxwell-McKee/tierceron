package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"

	"github.com/trimble-oss/tierceron-core/v2/core"
	pb "github.com/trimble-oss/tierceron/installation/trcshk/trchelloworld/hellosdk" // Update package path as needed
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type server struct {
	pb.UnimplementedGreeterServer
}

var configContext *ConfigContext
var grpcServer *grpc.Server
var sender chan error

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

const (
	HELLO_CERT  = "Common/hello.crt.mf.tmpl"
	HELLO_KEY   = "Common/hellokey.key.mf.tmpl"
	COMMON_PATH = "config"
)

func GetConfigPaths() []string {
	return []string{
		COMMON_PATH,
		HELLO_CERT,
		HELLO_KEY,
	}
}

func receiver(receive_chan chan int) {
	for {
		event := <-receive_chan
		switch {
		case event == core.PLUGIN_EVENT_START:
			go start()
		case event == core.PLUGIN_EVENT_STOP:
			go stop()
			sender <- errors.New("hello shutting down")
			return
		case event == core.PLUGIN_EVENT_STATUS:
			//TODO
		default:
			//TODO
		}
	}
}

func InitServer(port int, certBytes []byte, keyBytes []byte) (net.Listener, *grpc.Server, error) {
	var err error

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		log.Fatalf("Couldn't construct key pair: %v", err)
	}
	creds := credentials.NewServerTLSFromCert(&cert)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return nil, nil, err
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	return lis, grpcServer, nil
}

type ConfigContext struct {
	Config *map[string]interface{}
	Start  func()
	Cert   []byte
	Key    []byte
	Log    *log.Logger
}

func start() {
	if configContext == nil {
		fmt.Println("no config context initialized for healthcheck")
		return
	}
	if portInterface, ok := (*configContext.Config)["grpc_server_port"]; ok {
		var helloPort int
		if port, ok := portInterface.(int); ok {
			helloPort = port
		} else {
			var err error
			helloPort, err = strconv.Atoi(portInterface.(string))
			if err != nil {
				configContext.Log.Printf("Failed to process server port: %v", err)
				if sender != nil {
					sender <- err
				}
				return
			}
		}

		fmt.Printf("Server listening on :%d\n", helloPort)
		lis, gServer, err := InitServer(helloPort, configContext.Cert, configContext.Key)
		if err != nil {
			configContext.Log.Printf("Failed to start server: %v", err)
			if sender != nil {
				sender <- err
			}
			return
		}
		configContext.Log.Println("Starting server")

		grpcServer = gServer
		grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
		pb.RegisterGreeterServer(grpcServer, &server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			configContext.Log.Println("Failed to serve:", err)
			if sender != nil {
				sender <- err
			}
			return
		}
	} else {
		configContext.Log.Println("Missing config: gprc_server_port")
		if sender != nil {
			sender <- errors.New("missing config: gprc_server_port")
		}
		return
	}
}

func stop() {
	if grpcServer == nil || configContext == nil {
		fmt.Println("no server initialized for hello")
		return
	}
	configContext.Log.Println("Stopping server")
	fmt.Println("Stopping server")
	grpcServer.Stop()
	fmt.Println("Stopped server")
	configContext = nil
	grpcServer = nil
	sender = nil
}

func Init(properties *map[string]interface{}) {
	if properties == nil {
		fmt.Println("Missing initialization components")
		return
	}
	var logger *log.Logger
	if _, ok := (*properties)["log"].(*log.Logger); ok {
		logger = (*properties)["log"].(*log.Logger)
	}
	var certbytes []byte
	var keybytes []byte
	var config_properties *map[string]interface{}
	if cert, ok := (*properties)[HELLO_CERT]; ok {
		certbytes = cert.([]byte)
	}
	if key, ok := (*properties)[HELLO_KEY]; ok {
		keybytes = key.([]byte)
	}
	if common, ok := (*properties)[COMMON_PATH]; ok {
		config_properties = common.(*map[string]interface{})
	} else {
		fmt.Println("Missing common config components")
		return
	}

	configContext = &ConfigContext{
		Config: config_properties,
		Start:  start,
		Cert:   certbytes,
		Key:    keybytes,
		Log:    logger,
	}

	if channels, ok := (*properties)[core.PLUGIN_EVENT_CHANNELS_MAP_KEY]; ok {
		if chans, ok := channels.(map[string]interface{}); ok {
			if rchan, ok := chans[core.PLUGIN_CHANNEL_EVENT_IN]; ok {
				if rc, ok := rchan.(chan int); ok && rc != nil {
					go receiver(rc)
				} else {
					configContext.Log.Println("Unsupported receiving channel passed into hello")
					return
				}
			} else {
				configContext.Log.Println("No receiving channel passed into hello")
				return
			}
			if schan, ok := chans[core.PLUGIN_CHANNEL_EVENT_OUT]; ok {
				if sc, ok := schan.(chan error); ok && sc != nil {
					sender = sc
				} else {
					configContext.Log.Println("Unsupported sending channel passed into hello")
					return
				}
			} else {
				configContext.Log.Println("No sending channel passed into hello")
				return
			}
		} else {
			configContext.Log.Println("No channels passed into hello")
			return
		}
	}
}

func main() {
	logFilePtr := flag.String("log", "./trchelloworld.log", "Output path for log file")
	flag.Parse()
	config := make(map[string]interface{})

	f, err := os.OpenFile(*logFilePtr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		os.Exit(-1)
	}
	logger := log.New(f, "[trchelloworld]", log.LstdFlags)
	config["log"] = logger

	data, err := os.ReadFile("config.yml")
	if err != nil {
		logger.Println("Error reading YAML file:", err)
		os.Exit(-1)
	}

	// Create an empty map for the YAML data
	var configCommon map[string]interface{}

	// Unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &configCommon)
	if err != nil {
		logger.Println("Error unmarshaling YAML:", err)
		os.Exit(-1)
	}
	config[COMMON_PATH] = &configCommon

	helloCertBytes, err := os.ReadFile("./hello.crt")
	if err != nil {
		log.Printf("Couldn't load cert: %v", err)
	}

	helloKeyBytes, err := os.ReadFile("./hellokey.key")
	if err != nil {
		log.Printf("Couldn't load key: %v", err)
	}
	config[HELLO_CERT] = helloCertBytes
	config[HELLO_KEY] = helloKeyBytes

	Init(&config)
	configContext.Start()
}