package logging

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var opt Option
	if opt.IsEmpty() {
		t.Log("success")
		return
	}
	t.Log("failed")
}

func TestNewLogger(t *testing.T) {
	const (
		msg = "msg"
	)

	opt := Option{
		LogLevel: LevelALL,
		Format:   FormatFMT,
		Ouput:    Console,
	}

	logger := NewLogger(opt)
	logger.Log("coba", "disini")
}

func TestLogLevel(t *testing.T) {
	t.Fail()
	const (
		msg = "msg"
	)

	opt := Option{
		LogLevel: LevelALL,
		Format:   FormatFMT,
		Ouput:    Console,
	}

	logger := NewLogger(opt)

	logger.Info("ini info log")
	logger.Debug("ini debug log")
	logger.Warn("ini warning log")
	logger.Error("ini Error log")
}
