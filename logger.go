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

func (l *Logger) log(level string, message string, params interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("%s [%s] %s params: %v", timestamp, level, message, params)
	l.output.Println(logMessage)
}

func (l *Logger) Info(message string, params interface{}) {
	l.log("INFO", message, params)
}

func (l *Logger) Error(message string, params interface{}) {
	l.log("ERROR", message, params)
}

func (l *Logger) Warning(message string, params interface{}) {
	l.log("WARNING", message, params)
}

func (l *Logger) Fatal(message string, params interface{}) {
	l.log("FATAL", message, params)
}
