package log

import (
	"os"
	"time"

	"github.com/go-kit/kit/log"
	uc "github.com/tawakhal/util/constanta"
)

type LogConfig struct {
	LogFormat uc.LogFormat
}

func Logger(cfg LogConfig) log.Logger {
	// TODO
	return nil
}

func StdLogger(cfg LogConfig) log.Logger {
	logger := log.NewLogfmtLogger(os.Stdout)

	switch cfg.LogFormat {
	case uc.LFContext:
		// TODO
	case uc.LFDebug:
		logger = logDebugInfo(logger)
		return logger
	case uc.LFValuer:
		// TODO
	}
	return logger
}

func logDebugInfo(logger log.Logger) log.Logger {
	baseTime := time.Now()
	mockTime := func() time.Time {
		baseTime = baseTime.Add(time.Second)
		return baseTime
	}

	logger = log.With(logger, "time", log.Timestamp(mockTime), "caller", log.DefaultCaller)
	return logger
}

func logContextual() {
	// TODO
}
