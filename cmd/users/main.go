package main

import (
	"fmt"
	"os"
	"io"

	"github.com/gin-gonic/gin"
	stdopentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/vmwarecloudadvocacy/user/internal/auth"
	"github.com/vmwarecloudadvocacy/user/internal/db"
	"github.com/vmwarecloudadvocacy/user/internal/service"
	"github.com/vmwarecloudadvocacy/user/pkg/logger"
)

const (
	dbName         = "acmefit"
	collectionName = "users"
)

func initJaeger(service string) (stdopentracing.Tracer, io.Closer) {
	
	// Uncomment the lines below, if sending traces directly to the collector
	//tracerIP := GetEnv("TRACER_HOST", "localhost")
	//tracerPort := GetEnv("TRACER_PORT", "14268")
	
	agentIP := db.GetEnv("JAEGER_AGENT_HOST", "localhost")
    agentPort := db.GetEnv("JAEGER_AGENT_PORT", "6831")

	logger.Logger.Infof("Sending traces to %s %s", agentIP, agentPort)

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
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

// This handles initiation of "gin" router. It also defines routes to various APIs
// Env variable USER_IP and USER_PORT should be used to set IP and PORT.
// For example: export USER_PORT=8086 will start the server on local IP at 0.0.0.0:8086
func handleRequest() {

	// Init Router
	router := gin.New()

	nonAuthGroup := router.Group("/")
	{
		nonAuthGroup.POST("/register", service.RegisterUser)
		nonAuthGroup.POST("/login", service.LoginUser)
		nonAuthGroup.POST("/refresh-token", service.RefreshAccessToken)
		nonAuthGroup.POST("/verify-token", service.VerifyAuthToken)
	}

	authGroup := router.Group("/")
	// Added
	authGroup.Use(auth.AuthMiddleware())
	{
		authGroup.GET("/users", service.GetUsers)
		authGroup.GET("/users/:id", service.GetUser)
		authGroup.DELETE("/users/:id", service.DeleteUser)
		authGroup.POST("/logout", service.LogoutUser)
	}

	//flag.Parse()

	// Set default values if ENV variables are not set
	port := db.GetEnv("USERS_PORT", "8081")
	ip := db.GetEnv("USERS_HOST", "0.0.0.0")

	ipPort := ip + ":" + port

	logger.Logger.Infof("Starting user service at %s on %s", ip, port)

	router.Run(ipPort)

}

func main() {

	//create your file with desired read/write permissions
	f, err := os.OpenFile("log.info", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Could not open file ", err)
		logger.Logger.Infof("Could not open file")
	} else {
		logger.InitLogger(f)
	}

	dbsession := db.ConnectDB(dbName, collectionName)
	logger.Logger.Infof("Successfully connected to database %s", dbName)

	redisClient := db.ConnectRedisDB()
	
	tracer, closer := initJaeger("user")

	stdopentracing.SetGlobalTracer(tracer)

	handleRequest()

	db.CloseDB(dbsession)

	defer closer.Close()

	defer f.Close()

	defer redisClient.Close()

}
