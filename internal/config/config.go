package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title   string
	Env     string
	Version string
}

func newDefaultConfig() *Config {
	return &Config{
		Title:   "XiaRen",
		Env:     "debug",
		Version: "0.0.0",
	}
}

func LoadConfig(filepath string) error {
	return sync.OnceValue(func() error {
		var loadedConfg = Config{}
		_, err := toml.DecodeFile(filepath, &loadedConfg)
		if err != nil {
			return err
		}
		_conf = &loadedConfg
		return nil
	})()
}

var _conf *Config

func GetConfig() *Config {
	if _conf == nil {
		_conf = newDefaultConfig()
	}
	return _conf
}
