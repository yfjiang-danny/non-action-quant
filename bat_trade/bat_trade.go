package bat_trade

import (
	"log"
	"time"

	"github.com/yzlq99/eastmoneyapi/client"
	"github.com/yzlq99/eastmoneyapi/model"
	"github.com/yzlq99/non-action-quant/config"
	"github.com/yzlq99/non-action-quant/utils"
)

type EmptyModel struct {
	Time    string `json:"Time"`
	Message string `json:"Message"`
}

type LogModel struct {
	Time string `json:"Time"`
	*model.SubmitBatTradeResult
}

// 每周 1-5 的 11：08 申购新股新债
type BatTrade struct {
	EmCli *client.EastMoneyClient
}

func (b *BatTrade) Spec() string {
	// return "8 11 * * 1-5"
	return config.GetConfig().BatTradeSpec
}

func (b *BatTrade) Run() {
	go b.newConvertibleBond()
	go b.newStock()
}

func (b *BatTrade) newConvertibleBond() {
	currentTime := time.Now()
	bonds, err := b.EmCli.GetNewConvertibleBondList()
	if err != nil {
		log.Panic(err)
	}

	if bonds == nil || len(bonds.Data) <= 0 {
		log.Print(EmptyModel{
			Time:    currentTime.Format("2006-1-2 15:4:5"),
			Message: "今天无新债申购",
		})
		return
	}

	res, err := b.EmCli.SubmitBatTrade(bonds.GetSubmitBatTradeParams())
	if err != nil {
		log.Panic(err)
	}
	log.Print(utils.ToJson(LogModel{
		Time:                 currentTime.Format("2006-1-2 15:4:5"),
		SubmitBatTradeResult: res,
	}))
}

func (b *BatTrade) newStock() {
	currentTime := time.Now()
	newStock, err := b.EmCli.GetCanBuyNewStockList()
	if err != nil {
		log.Panic(err)
	}

	if newStock == nil || len(newStock.NewStockList) <= 0 {
		log.Print(EmptyModel{
			Time:    currentTime.Format("2006-1-2 15:4:5"),
			Message: "今天无新债申购",
		})
		return
	}

	res, err := b.EmCli.SubmitBatTrade(newStock.GetSubmitBatTradeParams())
	if err != nil {
		log.Panic(err)
	}
	log.Print(utils.ToJson(LogModel{
		Time:                 currentTime.Format("2006-1-2 15:4:5"),
		SubmitBatTradeResult: res,
	}))
}
