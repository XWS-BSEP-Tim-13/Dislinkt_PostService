package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Logger struct {
	Logger *logrus.Logger
	file   *os.File
}

func InitLogger(logFile string, ctx context.Context) *Logger {
	logger := &Logger{}

	path := filepath.Join("logs", logFile+".log")
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	logger.Logger = logrus.New()
	logger.Logger.SetOutput(file)
	logger.Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02t15:04:05",
		FullTimestamp:   true,
	})
	logger.Logger.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: logger.Logger, file: logger.file}
}

func (l *Logger) WarningMessage(message string) {
	l.Logger.Warning(message)
}

func (l *Logger) ErrorMessage(message string) {
	l.Logger.Error(message)
}

func (l *Logger) InfoMessage(message string) {
	l.Logger.Info(message)
}

func (l *Logger) FatalMessage(message string) {
	l.Logger.Fatal(message)
}
