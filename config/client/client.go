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

	// Testing String
	satuString := cfg.AppData.Get("string", "malah coba dong")
	arrayString := cfg.AppData.GetA("stringarray", []string{"malah,coba ,coba"})
	devSatString := cfg.AppData.Get("devSatString", "GA akan ADA")
	devArrString := cfg.AppData.GetA("devArrString", []string{"MASA , NGAK , ADA "})
	logger.Log("info", fmt.Sprintf("%v", satuString))
	logger.Log("info", fmt.Sprintf("%v", arrayString))
	logger.Log("info", fmt.Sprintf("%v", devSatString))
	logger.Log("info", fmt.Sprintf("%v", devArrString))

	satuInt := cfg.AppData.GetInt("int", 1)
	arrayInt := cfg.AppData.GetAInt("intarray", []int{0, 9, 8, 7, 6, 5})
	devSatInt := cfg.AppData.GetInt("devSatInt", 61782)
	devArrInt := cfg.AppData.GetAInt("devArrInt", []int{7, 1, 5, 2, 3})
	logger.Log("info", fmt.Sprintf("%v", satuInt))
	logger.Log("info", fmt.Sprintf("%v", arrayInt))
	logger.Log("info", fmt.Sprintf("%v", devSatInt))
	logger.Log("info", fmt.Sprintf("%v", devArrInt))
}
