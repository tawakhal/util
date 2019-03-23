package main

import (
	uc "github.com/tawakhal/util/constanta"
	"github.com/tawakhal/util/log"
)

func main() {
	cfg := log.LogConfig{LogFormat: uc.LFDebug}
	logger := log.Logger(cfg)
	logger.Log("Testing", "Olgi Tawakhal")
}
