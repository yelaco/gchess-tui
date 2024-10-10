package util

import (
	"net/url"

	"github.com/spf13/viper"
)

type Env struct {
	RunMode string `mapstructure:"RUN_MODE"`
}

// Config stores all configurations of the application
type Config struct {
	Debug      bool     `mapstructure:"debug"`
	ServiceUrl *url.URL `mapstructure:"SERVICE_URL"`
}

// LoadConfig reads configurations from file or environment variables
func LoadEnv() (env Env, err error) {
	viper.SetConfigFile(".env")
	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&env); err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}

func LoadConfig(path string) (config Config, err error) {
	env, err := LoadEnv()
	if err != nil {
		return
	}

	viper.SetDefault("debug", true)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	switch env.RunMode {
	case "dev":
		viper.SetConfigName("dev")
	case "prod":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("dev")
	}

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
