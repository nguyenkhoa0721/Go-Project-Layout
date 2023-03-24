package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database Database
	Redis    Redis
	Uuid     Uuid
	Kafka    Kafka
	Mongo    Mongo
	Build    Build
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		config = nil
		return config, err
	}

	return config, nil
}
