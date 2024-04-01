package logger

import (
	"errors"
	"log"
	"os"
)

var (
	ErrorFlag = 0
	WarnFlag  = 1
	InfoFlag  = 2
)

type Logger struct {
	file        *os.File
	errorLogger *log.Logger
	warnLogger  *log.Logger
	infoLogger  *log.Logger
	Debug       bool
}

var (
	ErrCantOpenOrCreateLogFile = errors.New("can't open or create log file")
)

func New(path string) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Logger{
		file:        file,
		errorLogger: log.New(file, "[ERROR]: ", ErrorFlag),
		warnLogger:  log.New(file, "[WARN]: ", WarnFlag),
		infoLogger:  log.New(file, "[INFO]: ", InfoFlag),
	}, nil
}

func (logger *Logger) Error(v ...any) {
	logger.errorLogger.Println(v...)
	if logger.Debug {
		log.Println(v...)
	}
}
func (logger *Logger) Errorf(format string, v ...any) {
	logger.errorLogger.Printf(format, v...)
	if logger.Debug {
		log.Printf(format, v...)
	}
}

func (logger *Logger) Warn(v ...any) {
	logger.warnLogger.Println(v...)
	if logger.Debug {
		log.Println(v...)
	}
}
func (logger *Logger) Warnf(format string, v ...any) {
	logger.warnLogger.Printf(format, v...)
	if logger.Debug {
		log.Printf(format, v...)
	}
}

func (logger *Logger) Info(v ...any) {
	logger.infoLogger.Println(v...)
	if logger.Debug {
		log.Println(v...)
	}
}
func (logger *Logger) Infof(format string, v ...any) {
	logger.infoLogger.Printf(format, v...)
	if logger.Debug {
		log.Printf(format, v...)
	}
}

func (logger *Logger) Close() error {
	return logger.file.Close()
}
