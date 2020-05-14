package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/alexander-molina/avito_task/internal/app/api"
	"github.com/alexander-molina/avito_task/internal/app/config"
)

var (
	subnetPrefixSize int
	reqestLimit      int
	blockTime        int
)

func init() {
	flag.IntVar(&subnetPrefixSize, "p", 24, "The size of subnet prefix the server will work with")
	flag.IntVar(&reqestLimit, "r", 100, "Quantity of maximum permitted requests per tme period")
	flag.IntVar(&blockTime, "b", 2, "Duration of block period (in minutes)")
	flag.Parse()

	appConfig := config.GetConfig()
	appConfig.SubnetPrefixSize = strconv.Itoa(subnetPrefixSize)
	appConfig.ReqestLimit = reqestLimit
	appConfig.BlockTime = time.Minute * time.Duration(blockTime)

	config.SetConfig(appConfig)
}

func main() {
	api.StartServer()
}
