package main

import (
	"log"
	"os"
)

import config "pocket-trader-backend/config"
import logger "pocket-trader-backend/logger"
import server "pocket-trader-backend/rpc-server"

const defaultConfigPath = "main-server.config.toml"

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

	go server.RunGrpc(c.Server.GrpcPort, c.Db.IndustriesSource)
	server.RunRest(c.Server.Ip, c.Server.GrpcPort, c.Server.HttpPort)
}
