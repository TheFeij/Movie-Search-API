package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config holds configuration values of the app
type Config struct {
	ServerAddress          string        `mapstructure:"SERVER_ADDRESS"`
	ElasticSearchAddress   string        `mapstructure:"ELASTIC_SEARCH_ADDRESS"`
	XRapidAPIKey           string        `mapstructure:"X_RAPID_API_KEY"`
	XRapidAPIHost          string        `mapstructure:"X_RAPID_API_HOST"`
	RapidAPISearchEndpoint string        `mapstructure:"RAPID_API_SEARCH_ENDPOINT"`
	RedisAddress           string        `mapstructure:"REDIS_ADDRESS"`
	RedisDialTimeout       time.Duration `mapstructure:"REDIS_DIAL_TIMEOUT"`
	RedisReadTimeout       time.Duration `mapstructure:"REDIS_READ_TIMEOUT"`
	RedisDB                int           `mapstructure:"REDIS_DB"`
}

// LoadConfig loads configuration values from config.env to config struct
func LoadConfig(path string, name string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config file:\n%v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config file:\n%v", err)
	}

	return config, nil
}
