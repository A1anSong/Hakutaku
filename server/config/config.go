package config

import (
	"github.com/spf13/viper"
)

var (
	Configs = Config{
		Log: LogConfig{
			Level: "info",
			File:  "server.log",
		},
	}
)

type LogConfig struct {
	Level string
	File  string
}

type DBConfig struct {
	Driver string
	Uri    string
}

type Config struct {
	Log LogConfig
	DB  DBConfig
}

func LoadConfig(cfg *viper.Viper) (err error) {
	err = cfg.Unmarshal(&Configs)
	return
}
