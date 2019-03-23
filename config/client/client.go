package main

import (
	"fmt"

	cfg "github.com/tawakhal/util/config"
	uc "github.com/tawakhal/util/constanta"
	"github.com/tawakhal/util/log"
)

func main() {
	logCfg := log.LogConfig{LogFormat: uc.LFDebug}
	logger := log.StdLogger(logCfg)
	err := cfg.AppData.LoadConfig()
	if err != nil {
		logger.Log("error", err.Error())
	}
	satuString := cfg.AppData.Get("string", "malah coba dong")
	arrayString := cfg.AppData.GetA("stringarray", []string{"malah,coba ,coba"})
	devSatString := cfg.AppData.Get("devSatString", "GA akan ADA")
	devArrString := cfg.AppData.GetA("devArrString", []string{"MASA , NGAK , ADA "})
	logger.Log("info", fmt.Sprintf("%v", satuString))
	logger.Log("info", fmt.Sprintf("%v", arrayString))
	logger.Log("info", fmt.Sprintf("%v", devSatString))
	logger.Log("info", fmt.Sprintf("%v", devArrString))
}
