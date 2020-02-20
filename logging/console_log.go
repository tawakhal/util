package logging

import (
	"log"
	"time"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/tawakhal/util/uuid"
)

type consoleLog struct {
	requestID uint64
	timeStart time.Time
	log       logkit.Logger
}

// NewConsoleLog is logger with console log
func NewConsoleLog(option Option) Logger {
	if option.IsEmpty() {
		log.Fatal("option is empty")
	}
	lg := option.LogLevel.newLevel(option.Format.newFormat(option.Ouput.newOutput(), 6))
	level.Info(lg).Log("range", option.LogLevel, "format", option.Format, "output", option.Ouput)

	nID, _ := uuid.New()
	return &consoleLog{
		requestID: nID.MSB,
		timeStart: time.Now(),
		log:       lg,
	}
}

// Log is default logging
func (cl *consoleLog) Log(keyvals ...interface{}) error {
	newVals := cl.defaultConsoleLog()
	newVals = append(newVals, keyvals...)
	return cl.log.Log(keyvals...)
}

// Info is logging with level info
func (cl *consoleLog) Info(keyvals ...interface{}) error {
	newVals := cl.defaultConsoleLog()
	newVals = append(newVals, keyvals...)
	return level.Info(cl.log).Log(newVals...)
}

// Debug is logging with level Debug
func (cl *consoleLog) Debug(keyvals ...interface{}) error {
	newVals := cl.defaultConsoleLog()
	newVals = append(newVals, keyvals...)
	return level.Debug(cl.log).Log(newVals...)
}

// Warn is logging with level Warn
func (cl *consoleLog) Warn(keyvals ...interface{}) error {
	newVals := cl.defaultConsoleLog()
	newVals = append(newVals, keyvals...)
	return level.Warn(cl.log).Log(newVals...)
}

// Error is logging with level Error
func (cl *consoleLog) Error(keyvals ...interface{}) error {
	newVals := cl.defaultConsoleLog()
	newVals = append(newVals, keyvals...)
	return level.Error(cl.log).Log(newVals...)
}

func (cl *consoleLog) defaultConsoleLog() []interface{} {
	var newVals []interface{}
	newVals = append(newVals, "req", cl.requestID, "since", time.Since(cl.timeStart))
	return newVals
}
