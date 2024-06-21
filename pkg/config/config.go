package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type config struct {
	WebDirectory string
}

var Config *config

func SetupConfig() *config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "failed to read in config"))
	}

	Config = &config{}

	if err := viper.Unmarshal(Config); err != nil {
		panic(errors.Wrap(err, "unable to decode into struct"))
	}

	return Config
}
