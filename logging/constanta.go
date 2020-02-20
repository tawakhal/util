package logging

import (
	"fmt"
	"io"
	"os"

	logkit "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

func (et FormatLog) newFormat(w io.Writer, call int) logkit.Logger {
	var logger logkit.Logger
	switch et {
	default:
		logger = logkit.NewLogfmtLogger(w)
	case FormatJSON:
		logger = logkit.NewJSONLogger(w)
	case FormatFMT:
		logger = logkit.NewLogfmtLogger(w)
	}
	if call == 0 {
		call = 5
	}
	return logkit.With(logger, "time", logkit.DefaultTimestampUTC, "caller", logkit.Caller(call))
}

// OutputLog is output of log. its to console or file
type OutputLog int32

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
