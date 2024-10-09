package util

import (
	"github.com/spf13/viper"
)

// Config stores all confdigurations of the application
// The values are read by viper from a config file or environment variables.
type Config struct {
	RunMode    string `mapstructure:"RUN_MODE"`
	ServiceUrl string `mapstructure:"SERVCIE_URL"`
}

// LoadConfig reads configurations from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
