package util

import (
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config stores all confdigurations of the application
// The values are read by viper from a config file or environment variables.
type Config struct {
	RunMode    string
	ServiceUrl *url.URL `mapstructure:"SERVICE_URL"`
}

// LoadConfig reads configurations from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	runMode := os.Getenv("RUN_MODE")
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	switch runMode {
	case "development":
		viper.SetConfigName("dev")
	case "stage":
	case "production":
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
