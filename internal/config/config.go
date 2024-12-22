package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/morheus9/go_rest/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"    env-required:"port"`
	Listen  struct {
		Type   string `yaml:"type"    env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port"    env-default:"8080"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	// Only once starting, singleton
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read aplication config")
		instance = &Config{}
		// Всё из OS
		// if err := cleanenv.ReadEnv()
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
