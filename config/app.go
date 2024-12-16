package config

import (
	"log"
	"os"
)

var AppLogger *log.Logger

func SetupLogger() {
	logFile, err := os.OpenFile(os.Getenv("LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	AppLogger = log.New(logFile, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("Logger initialized")
}
