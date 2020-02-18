package logging

import (
	"fmt"
	"io"
	"log"
	"os"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Level is list of level
type Level int32

// List of const Level
const (
	LevelALL   Level = 1
	LevelErr   Level = 2
	LevelDebug Level = 3
	LevelInfo  Level = 4
	LevelWarn  Level = 5
)

// Int32 is convert from type to int32
func (et Level) Int32() int32 {
	return int32(et)
}

// String is convert from type to string
func (et Level) String() string {
	switch et {
	case LevelALL:
		return "ALL Level Log"
	case LevelErr:
		return "Error Level"
	case LevelDebug:
		return "Debug Level"
	case LevelInfo:
		return "Info Level"
	case LevelWarn:
		return "Warning Level"
	default:
		return fmt.Sprintf("miss match : %d", et)
	}
}

func (et Level) newLevel(logger logkit.Logger) logkit.Logger {
	var opt level.Option
	switch et {
	default:
		opt = level.AllowInfo()
	case LevelALL:
		opt = level.AllowAll()
	case LevelErr:
		opt = level.AllowError()
	case LevelDebug:
		opt = level.AllowDebug()
	case LevelInfo:
		opt = level.AllowInfo()
	case LevelWarn:
		opt = level.AllowWarn()
	}

	logger = level.NewFilter(logger, opt)
	return logger
}

// FormatLog is list of format log
type FormatLog int32

// List of Format Log
const (
	FormatJSON FormatLog = 1
	FormatFMT  FormatLog = 2
)

// String is convert from type to string
func (et FormatLog) String() string {
	switch et {
	case FormatJSON:
		return "JSON Format"
	case FormatFMT:
		return "FMT Format"
	default:
		return fmt.Sprintf("miss match : %d", et)
	}
}

// Int32 is convert from type to int32
func (et FormatLog) Int32() int32 {
	return int32(et)
}

func (et FormatLog) newFormat(w io.Writer) logkit.Logger {
	switch et {
	default:
		return logkit.NewLogfmtLogger(w)
	case FormatJSON:
		return logkit.NewJSONLogger(w)
	case FormatFMT:
		return logkit.NewLogfmtLogger(w)
	}
}

// OutputLog is output of log. its to console or file
type OutputLog int32

// Option is option for logging
type Option struct {
	LogLevel Level
	Format   FormatLog
	Ouput    OutputLog
}

// IsEmpty is checking struct option is empty or not
func (et Option) IsEmpty() bool {
	var opt Option
	if et == opt {
		return true
	}
	return false
}

// List of Format Log
const (
	Console OutputLog = 1
	File    OutputLog = 2
)

// String is convert from type to string
func (et OutputLog) String() string {
	switch et {
	case Console:
		return "Output to console"
	case File:
		return "Output to file"
	default:
		return fmt.Sprintf("miss match : %d", et)
	}
}

// Int32 is convert from type to int32
func (et OutputLog) Int32() int32 {
	return int32(et)
}
func (et OutputLog) newOutput() io.Writer {
	const (
		fileName = "service.log"
	)
	switch et {
	default:
		file(fileName)
		return NewDefaultLogWriter()
	case File:
		file(fileName)
		return NewDefaultLogWriter()
	case Console:
		return os.Stderr
	}
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

// NewLogger is create logger with option
func NewLogger(option Option) logkit.Logger {
	if option.IsEmpty() {
		log.Fatal("option is empty")
	}
	logger := option.LogLevel.newLevel(option.Format.newFormat(option.Ouput.newOutput()))
	level.Info(logger).Log("range", option.LogLevel, "format", option.Format, "output", option.Ouput)
	return logger
}
