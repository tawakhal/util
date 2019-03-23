package log

import (
	"testing"

	uc "github.com/tawakhal/util/constanta"
)

func logDebug(t *testing.T) {
	cfg := LogConfig{LogFormat: uc.LFDebug}
	logger := Logger(cfg)
	logger.Log("Testing", "Olgi Tawakhal")
}
