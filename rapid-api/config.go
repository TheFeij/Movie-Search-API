package rapid_api

import (
	"fmt"
	"github.com/spf13/viper"
)

var config struct {
	XRapidAPIKey           string `mapstructure:"X_RAPID_API_KEY"`
	XRapidAPIHost          string `mapstructure:"X_RAPID_API_HOST"`
	RapidAPISearchEndpoint string `mapstructure:"RAPID_API_SEARCH_ENDPOINT"`
}

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