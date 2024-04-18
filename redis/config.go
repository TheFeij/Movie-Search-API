package redis

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// config holds configuration values for this package
var config struct {
	RedisAddress     string        `mapstructure:"REDIS_ADDRESS"`
	RedisDialTimeout time.Duration `mapstructure:"REDIS_DIAL_TIMEOUT"`
	RedisReadTimeout time.Duration `mapstructure:"REDIS_READ_TIMEOUT"`
	RedisDB          int           `mapstructure:"REDIS_DB"`
}

// loadConfig loads configuration values from config.json to config struct
func loadConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file:\n%v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal config file:\n%v", err)
	}

	return nil
}
