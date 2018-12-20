package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type service struct {
	HOST string `mapstructure:"host"`
	PORT string `mapstructure:"port"`
}

type db struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

// Config config struct
var Config appConfig

type appConfig struct {
	MODE    string `mapstructure:"mode"`
	SERVICE service
	DB      db
}

// LoadConfig load configs
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("app")
	v.AutomaticEnv()

	// set default configs

	for _, configPath := range configPaths {
		v.AddConfigPath(configPath)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
