package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Ip            string `mapstructure:"ip"`
		Port          int    `mapstructure:"port"`
		TickerTimeMin int    `mapstructure:"ticker_time_min"`
	} `mapstructure:"main-server"`
	LogPath string `mapstructure:"log_path"`
}

func ReadConfig(filepath string) *Config {
	config := Config{}
	viper.SetConfigFile(filepath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}
	return &config
}

func (c Config) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("error printing config: ", err)
	}
	return "config: " + string(b)
}
