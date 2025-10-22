package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

// function sets up logging to both console and writes to file (if available)
func Init() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[WARN] Could not open log file, logging to console only")
		file = os.Stdout //falls back to console logging
	}

	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(file, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

}

func Info(message string) {

	log.Printf("[INFO] %s", message)

}

func Error(err error, context string) {

	log.Printf("[ERROR] %s: %v", context, err)

}

func Warn(message string) {

	log.Printf("[WARN] %s", message)

}
