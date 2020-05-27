package logger

import (
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once sync.Once

	//Logger is the global logging wrapper
	logger *logrus.Logger
)

// InitLogger initialises the global logger
func InitLogger() {
	once.Do(func() {
		logger = logrus.New()
		logger.Level = logrus.DebugLevel
	})
}

// SetOutput changes the output for the logger
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

// Debug logs a message at level Debug on the global logger.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf logs a message at level Debug on the global logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Debugln logs a message at level Debug on the global logger.
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Info logs a message at level Info on the global logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs a message at level Info on the global logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Infoln logs a message at level Info on the global logger.
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Error logs a message at level Error on the global logger.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf logs a message at level Error on the global logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Errorln logs a message at level Error on the global logger.
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}
