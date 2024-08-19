package utils

import (
	"log"
	"os"
)

var errorLog = log.New(os.Stderr, "go-git: Error: ", log.Ldate|log.Ltime|log.Lshortfile)
var commonLog = log.New(os.Stdout, "go-git: Log", log.Ldate|log.Ltime|log.Lshortfile)

func ErrorLogger(args ...interface{}) {
	errorLog.Println(args...)
}

func ErrorLoggerF(s string, args ...interface{}) {
	errorLog.Printf(s, args...)
}

func CommonLogger(args ...interface{}) {
	commonLog.Println(args...)
}

func CommonLoggerF(s string, args ...interface{}) {
	commonLog.Printf(s, args...)
}
