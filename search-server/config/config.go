package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	configPath = "./config.json"
	logFile    = "./log.txt"
)

var (
	DefaultConfig conf
)

type Log struct {
	Output string `json:"output"`
	Level  string `json:"level"`
}

type conf struct {
	HTTPAddress        string `json:"http_address"`
	SearchEngineDBPath string `json:"search_engine_db_path"`
	Log                Log `json:"log"`
}

func init() {
	f, err := os.Open(configPath)
	if nil != err {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&DefaultConfig)
	if nil != err {
		panic(err)
	}

	if "file" == DefaultConfig.Log.Output {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0666)
		if nil != err {
			panic(err)
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(os.Stdout)
	}

	level, err := log.ParseLevel(DefaultConfig.Log.Level)
	if nil != err {
		panic(err)
	}

	log.SetLevel(level)
}
