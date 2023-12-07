package main

import (
	"log"
	"os"
	config "pocket-trader-backend/config"
	logger "pocket-trader-backend/logger"
)

const defaultConfigPath = "server.config.toml"

func main() {
	configPath, exists := os.LookupEnv("CONFIG")
	if !exists {
		configPath = defaultConfigPath
	}

	c := config.ReadConfig(configPath)

	logPath := c.LogPath
	if len(logPath) != 0 {
		logger.Init(logger.FileMode, logPath)
	} else {
		logger.Init(logger.ConsoleMode)
	}

	log.Println(c)
}
