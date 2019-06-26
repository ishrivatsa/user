package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	logger *logrus.Logger
	zip    = flag.String("zipkin", os.Getenv("ZIPKIN"), "Zipkin address")
	//	port        = flag.String("port", "8080", "Port number on which the service should run")
	//	ip          = flag.String("ip", "0.0.0.0", "Preferred IP address to run the service on")
	serviceName = "user"
)

const (
	dbName         = "acmefit"
	collectionName = "users"
)

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	
	// Uncomment the lines below, if sending traces directly to the collector
	//tracerIP := GetEnv("TRACER_HOST", "localhost")
	//tracerPort := GetEnv("TRACER_PORT", "14268")
	
	agentIP := GetEnv("JAEGER_AGENT_HOST", "localhost")
    agentPort := GetEnv("JAEGER_AGENT_PORT", "6831")

	logger.Infof("Sending traces to %s %s", tracerIP, tracerPort)

	cfg := &jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:          true,
			LocalAgentHostPort: agentIP + ":" + agentPort,
// Uncomment the line below, if sending traces directly to the collector
//			CollectorEndpoint: "http://" + tracerIP + ":" + tracerPort + "/api/traces",
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

// GetEnv accepts the ENV as key and a default string
// If the lookup returns false then it uses the default string else it leverages the value set in ENV variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	logger.Info("Setting default values for ENV variable " + key)
	return fallback
}

// This initiates a new Logger and defines the format for logs
func initLogger(f *os.File) {

	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "",
		PrettyPrint:     true,
	})

	// Set output of logs to Stdout
	// Change to f for redirecting to file
	logger.SetOutput(os.Stdout)

}

// This handles initiation of "gin" router. It also defines routes to various APIs
// Env variable USER_IP and USER_PORT should be used to set IP and PORT.
// For example: export USER_PORT=8086 will start the server on local IP at 0.0.0.0:8086
func handleRequest() {

	// Init Router
	router := gin.Default()

	// Added
	authGroup := router.Group("/")
	{
		authGroup.POST("/register", RegisterUser)
		authGroup.POST("/login", LoginUser)
		authGroup.GET("/users", GetUsers)
		authGroup.GET("/users/:id", GetUser)
		authGroup.DELETE("/users/:id", DeleteUser)
	}

	//flag.Parse()

	// Set default values if ENV variables are not set
	port := GetEnv("USERS_PORT", "8081")
	ip := GetEnv("USERS_HOST", "0.0.0.0")

	ipPort := ip + ":" + port

	logger.Infof("Starting user service at %s on %s", ip, port)

	router.Run(ipPort)

}

func main() {

	//create your file with desired read/write permissions
	f, err := os.OpenFile("log.info", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Could not open file ", err)
	} else {
		initLogger(f)
	}

	dbsession := ConnectDB(dbName, collectionName, logger)

	logger.Infof("Successfully connected to database %s", dbName)

	tracer, closer := initJaeger("user")

	stdopentracing.SetGlobalTracer(tracer)

	handleRequest()

	CloseDB(dbsession, logger)

	defer closer.Close()

	defer f.Close()

}
