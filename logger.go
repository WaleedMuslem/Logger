package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	output *log.Logger
}

func NewLogger(filePath string) (*Logger, error) {
	var logger Logger
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		logger.output = log.New(file, "", 0)
	} else {
		logger.output = log.New(os.Stdout, "", 0)
	}
	return &logger, nil
}

func (l *Logger) log(level string, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("%s [%s] %s ", timestamp, level, message)
	l.output.Println(logMessage)
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

func (l *Logger) Fatal(message string) {
	l.log("FATAL", message)
}
