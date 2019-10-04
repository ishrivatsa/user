package logger

import (
	"flag"
	"os"
	"github.com/sirupsen/logrus"
)
var (
	Logger *logrus.Logger
	zip    = flag.String("zipkin", os.Getenv("ZIPKIN"), "Zipkin address")
	//	port        = flag.String("port", "8080", "Port number on which the service should run")
	//	ip          = flag.String("ip", "0.0.0.0", "Preferred IP address to run the service on")
	serviceName = "user"
)

// This initiates a new Logger and defines the format for logs
func InitLogger(f *os.File) {

	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "",
		PrettyPrint:     true,
	})

	// Set output of logs to Stdout
	// Change to f for redirecting to file
	Logger.SetOutput(os.Stdout)

}