package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Logger struct {
	InfoLogger  *logrus.Logger
	WarnLogger  *logrus.Logger
	ErrorLogger *logrus.Logger
}

func InitLogger(logFile string, ctx context.Context) *Logger {
	logger := &Logger{}
	logger.InfoLogger = InitLoggerPerLevel(logFile + "-info")
	logger.WarnLogger = InitLoggerPerLevel(logFile + "-warn")
	logger.ErrorLogger = InitLoggerPerLevel(logFile + "-error")

	logger.InfoLogger.SetLevel(logrus.InfoLevel)
	logger.WarnLogger.SetLevel(logrus.WarnLevel)
	logger.ErrorLogger.SetLevel(logrus.ErrorLevel)

	return &Logger{InfoLogger: logger.InfoLogger, WarnLogger: logger.WarnLogger, ErrorLogger: logger.ErrorLogger}
}

func InitLoggerPerLevel(logFile string) *logrus.Logger {
	path := filepath.Join("logs", logFile+".log")
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02t15:04:05",
		FullTimestamp:   true,
	})

	return logger
}

func (l *Logger) InfoMessage(message string) {
	l.InfoLogger.Info(message)
}

func (l *Logger) WarningMessage(message string) {
	l.WarnLogger.Warning(message)
}

func (l *Logger) ErrorMessage(message string) {
	l.ErrorLogger.Error(message)
}
