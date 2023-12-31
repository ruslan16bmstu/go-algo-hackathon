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
		GrpcPort      int    `mapstructure:"grpc_port"`
		HttpPort      int    `mapstructure:"http_port"`
		TickerTimeMin int    `mapstructure:"ticker_time_min"`
	} `mapstructure:"server"`
	Db struct {
		IndustriesSource  string `mapstructure:"industries_source"`
		PredictionsSource string `mapstructure:"predictions_source"`
	} `mapstructure:"db"`
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
