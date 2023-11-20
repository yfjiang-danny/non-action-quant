package main

import (
	"flag"
	"log"

	"github.com/yfjiang-danny/eastmoneyapi/client"
	"github.com/yfjiang-danny/non-action-quant/bat_trade"
	"github.com/yfjiang-danny/non-action-quant/config"
	"github.com/yfjiang-danny/non-action-quant/cron"
)

var configPath string

func init() {
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&configPath, "config", "", "")
	flag.Parse()
	if configPath != "" {
		config.SetConfigPath(configPath)
	}
}

func main() {

	emCli := client.NewEastMoneyClient(config.GetConfig().EastMoneyClientConfig)

	// 创建一个cron调度器
	cron.InitCron()

	cron.CronTab.AddJob(&bat_trade.BatTrade{
		EmCli: emCli,
	})

	cron.CronTab.Run()
}
