package config

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
)

// "/usr/local/echoserver/"
const (
	defaultAddr        = ":8001"
	defaultLogFilePath = "bin/echoserver.log"
	defaultConfigFile  = "bin/echoserver.json"
)

type Config struct {
	Addr        string `json:"addr"`
	LogFilePath string `json:"logFilePath"`
}

var configFile string
var Conf = &Config{
	Addr:        defaultAddr,
	LogFilePath: defaultLogFilePath,
}

func Init() error {

	flag.StringVar(&configFile, "config", defaultConfigFile, "")
	flag.Parse()

	data, err := os.ReadFile(configFile)
	if err != nil {
		if err := createDefaultConfigFile(); err != nil {
			return err
		} else {
			return nil
		}
	}

	if err := json.Unmarshal(data, Conf); err != nil {
		return err
	}

	return nil
}

func createDefaultConfigFile() error {

	data, err := json.Marshal(Conf)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(defaultConfigFile), 0755); err != nil {
		return err
	}

	return os.WriteFile(defaultConfigFile, data, 0644)
}
