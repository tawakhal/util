package logging

import (
	"log"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is interface of logger
type Logger interface {
	Log(keyvals ...interface{}) error
	Info(keyvals ...interface{}) error
	Debug(keyvals ...interface{}) error
	Warn(keyvals ...interface{}) error
	Error(keyvals ...interface{}) error
}

type logger struct {
	log logkit.Logger
}

// NewLogger is create logger with option
func NewLogger(option Option) Logger {
	if option.IsEmpty() {
		log.Fatal("option is empty")
	}
	lg := option.LogLevel.newLevel(option.Format.newFormat(option.Ouput.newOutput()))
	level.Info(lg).Log("range", option.LogLevel, "format", option.Format, "output", option.Ouput)
	return &logger{log: lg}
}

// Log is default logging
func (lg *logger) Log(keyvals ...interface{}) error {
	return lg.log.Log(checkKeyvals(keyvals))
}

// Info is logging with level info
func (lg *logger) Info(keyvals ...interface{}) error {

	return level.Info(lg.log).Log(checkKeyvals(keyvals))
}

// Debug is logging with level Debug
func (lg *logger) Debug(keyvals ...interface{}) error {
	return level.Debug(lg.log).Log(checkKeyvals(keyvals))
}

// Warn is logging with level Warn
func (lg *logger) Warn(keyvals ...interface{}) error {
	return level.Warn(lg.log).Log(checkKeyvals(keyvals))
}

// Error is logging with level Error
func (lg *logger) Error(keyvals ...interface{}) error {
	return level.Error(lg.log).Log(checkKeyvals(keyvals))
}

func checkKeyvals(keyvals []interface{}) []interface{} {
	if len(keyvals) != 0 && len(keyvals)%2 != 0 {
		keyvals = append(keyvals, keyvals[0])
	}
	return keyvals
}

// file set default log to file
func file(file string) {
	logFile := &lumberjack.Logger{
		Filename:  file,
		MaxSize:   1, // megabytes
		LocalTime: true,
		Compress:  true, // disabled by default
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
}
