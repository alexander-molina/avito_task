package main

import (
	"os"
	"strconv"
	"time"

	"github.com/alexander-molina/avito_task/internal/app/api"
	"github.com/alexander-molina/avito_task/internal/app/config"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8000"
	}

	subnetPrefixSize, exists := os.LookupEnv("PREFIX_SIZE")
	if !exists {
		subnetPrefixSize = "24"
	}

	reqestLimit, exists := os.LookupEnv("REQUEST_LIMIT")
	if !exists {
		reqestLimit = "100"
	}

	blockTime, exists := os.LookupEnv("BLOCK_TIME")
	if !exists {
		blockTime = "2"
	}

	appConfig := config.GetConfig()

	appConfig.SubnetPrefixSize = subnetPrefixSize
	appConfig.ReqestLimit, _ = strconv.Atoi(reqestLimit)
	t, _ := time.ParseDuration(blockTime + "m")
	appConfig.BlockTime = time.Minute * t

	config.SetConfig(appConfig)

	api.StartServer(":" + port)
}
