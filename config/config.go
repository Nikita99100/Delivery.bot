package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

const (
	envPrefix = "courierbot"
)

type Application struct {
	LogLevel      string
	LogFormat     string
	Port          string
	TelegramToken string
	CouriersUrl   string
}
type Config struct {
	Application Application
}

func Parse(filepath string) (*Config, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	bindEnvVars() // remember to parse the environment variables

	// Unmarshal the config
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("Application.loglevel", "debug")
	viper.SetDefault("Application.LogFormat", "text")
	viper.SetDefault("Application.port", "8080")
	viper.SetDefault("Application.telegramtoken", "tokenhere")
	viper.SetDefault("Application.couriersurl", "url")
}
func bindEnvVars() {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
