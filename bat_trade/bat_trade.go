package bat_trade

import (
	"log"
	"time"

	"github.com/yfjiang-danny/eastmoneyapi/client"
	"github.com/yfjiang-danny/eastmoneyapi/model"
	"github.com/yfjiang-danny/non-action-quant/config"
	"github.com/yfjiang-danny/non-action-quant/utils"
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
		log.Print(utils.ToJson(EmptyModel{
			Time:    currentTime.Format("2006-01-02 15:04:05"),
			Message: "今天无新债申购",
		}))
		return
	}

	res, err := b.EmCli.SubmitBatTrade(bonds.GetSubmitBatTradeParams())
	if err != nil {
		log.Panic(err)
	}
	log.Print(utils.ToJson(LogModel{
		Time:                 currentTime.Format("2006-01-02 15:04:05"),
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
		log.Print(utils.ToJson(EmptyModel{
			Time:    currentTime.Format("2006-01-02 15:04:05"),
			Message: "今天无新股申购",
		}))
		return
	}

	res, err := b.EmCli.SubmitBatTrade(newStock.GetSubmitBatTradeParams())
	if err != nil {
		log.Panic(err)
	}
	log.Print(utils.ToJson(LogModel{
		Time:                 currentTime.Format("2006-01-02 15:04:05"),
		SubmitBatTradeResult: res,
	}))
}
