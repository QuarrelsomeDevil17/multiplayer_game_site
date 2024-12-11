package utils

import (
	"log"
	"os"
)

// InitializeLogger sets up the logger with standard output and custom flags.
// Logs will include the date, time, and source file details.
func InitializeLogger() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Failed to log to file, using default stderr: %v", err)
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(logFile)
	}
	
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Logger initialized")
}
