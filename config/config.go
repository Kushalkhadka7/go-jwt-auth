package config

import (
	"time"

	"jwt-auth/constants"

	"github.com/spf13/viper"
)

// Config holds all the configurations to run api.
type Config struct {
	Server   *serverConfig
	DBConfig *DBConfig
}

// serverConfig holds all the configurations related to configure a server.
type serverConfig struct {
	Port         int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type DBConfig struct {
	Port         int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

// New initializes configurations for app.
func New() (*Config, error) {
	var c Config

	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(constants.ConfigFileFormat)
	viper.AddConfigPath(constants.ConfigFilePath)
	viper.AddConfigPath(constants.AddConfigPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	server := &serverConfig{
		Port:         c.Server.Port,
		ReadTimeOut:  c.Server.ReadTimeOut * time.Second,
		WriteTimeOut: c.Server.WriteTimeOut * time.Second,
	}

	dbConfig := &DBConfig{
		Port:         c.Server.Port,
		ReadTimeOut:  c.Server.ReadTimeOut * time.Second,
		WriteTimeOut: c.Server.WriteTimeOut * time.Second,
	}

	return &Config{Server: server, DBConfig: dbConfig}, nil
}
