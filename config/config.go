package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config map[string]string

const (
	sep = ","
)

var AppData Config
var configFile = "service.conf"

func (c *Config) Get(key, defval string) string {
	v, ok := AppData[key]
	if !ok {
		if defval != "" {
			AppData[key] = defval
			v = defval
		}
	}
	return v
}

func (c *Config) GetA(key string, defval []string) []string {
	v := c.Get(key, "")
	if len(v) == 0 {
		return defval
	}
	return strings.Split(v, sep)
}

func (c *Config) GetInt(key string, defval int) int {
	val := defval
	v := c.Get(key, "")
	if v != "" {
		val, _ = strconv.Atoi(v)
	}
	return val
}

func (c *Config) GetAInt(key string, defval []int) []int {
	val := defval
	v := c.Get(key, "")
	if v != "" {
		strArr := strings.Split(v, sep)
		val = []int{}
		for _, k := range strArr {
			kav, _ := strconv.Atoi(k)
			val = append(val, kav)
		}
	}
	return val
}

// LoadConfigDefault will generated config file with name service.conf in root folder
// ex :
// servicename/root_here
//			  /service.conf
func LoadConfigDefault() (*Config, error) {
	// TODO
	return nil, nil
}

// LoadConfig will read config name service.conf in root folder
func (c *Config) LoadConfig() error {
	err := c.open(configFile)
	return err
}

// LoadConfigFile will read config file from path file
// ex : file config at "/home/olgi/config/file.conf"
// so you can put path to parameter
func (c *Config) LoadConfigFile(file string) error {
	var err error
	if file != "" {
		err = c.open(file)
	} else {
		err = c.LoadConfig()
	}
	return err
}

func generatedConfigFile() {

}

func (c *Config) open(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	cfrdl, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(cfrdl, &c); err != nil {
		return err
	}
	return nil
}
