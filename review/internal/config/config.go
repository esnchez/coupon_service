package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	fileType = "yml"
	fileName = "config"
	filePath = "."
	portKey = "server.port"
	defaultPort = "8080"
	envTag = "SERVER_PORT"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
}

func Load() (*Config) {
	var config *Config

	viper := viper.New()
	viper.SetConfigType(fileType)
	viper.SetConfigName(fileName)
	viper.AddConfigPath(filePath)
	viper.AutomaticEnv()

	viper.SetDefault(portKey, defaultPort)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error loading from config file: %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("error marshaling from config: %s", err)
	}

	if err := viper.BindEnv(portKey, envTag); err != nil {
		log.Printf("error binding from env var: %s", err)
	}

	if config.Server.Port == 0 {
		config.Server.Port = viper.GetInt(portKey)
	}

	return config
}
