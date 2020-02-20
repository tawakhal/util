package logging

import "testing"

func TestNewConsoleLog(t *testing.T) {
	t.Fail()

	opt := Option{
		LogLevel: LevelALL,
		Format:   FormatFMT,
		Ouput:    Console,
	}

	logger := NewConsoleLog(opt)

	logger.Info("ini info log")
	logger.Debug("ini debug log")
	logger.Warn("ini warning log")
	logger.Error("ini Error log")

}
