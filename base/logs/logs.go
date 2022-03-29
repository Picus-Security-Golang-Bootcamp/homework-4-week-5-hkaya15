package logs

import (
	"log"
	"os"
)

func CreateLogs() *BuiltinLogger {	
	result := NewBuiltinLogger()
	return result
}

type BuiltinLogger struct {
	File   *os.File
	flag   int
	Logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	// os.0_APPEND append data to the file when writing.
	// os.0_CREATE create a new file if none exists.
	// os.0_WRONLY open the file write-only
	f, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return &BuiltinLogger{
		File:   f,
		flag:   log.LstdFlags,
		Logger: log.New(f, "LOGS: ", log.LstdFlags)}
}

