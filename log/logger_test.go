package logging

import (
	"testing"

	"github.com/go-kit/kit/log/level"
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

	level.Info(logger).Log(msg, "ini info log")
	level.Debug(logger).Log(msg, "ini debug log")
	level.Warn(logger).Log(msg, "ini warning log")
	level.Error(logger).Log(msg, "ini Error log")
}
