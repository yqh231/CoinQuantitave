package trx

import (
	"fmt"
	"testing"
)

func TestGetMarketList(t *testing.T) {
	coinMarket := CoinexMarket{}

	resp := coinMarket.GetMarketList()
	fmt.Println(*resp)
}


func TestGetMarketTicker(t *testing.T) {
	coinMarket := CoinexMarket{}

	resp := coinMarket.GetMarketTicker("BTCBCH")
	fmt.Println(*resp)
}

func TestGetMarketTickerAll(t *testing.T) {
	coinexMarket := CoinexMarket{}
	
	resp := coinexMarket.GetMarketTickerAll()
	fmt.Println(*resp)
}

func TestGetMarketDepth(t *testing.T) {
	coinMarket := CoinexMarket{}

	resp := coinMarket.GetMarketDepth("BTCBCH", "0", "5")
	fmt.Println(*resp)
}

func TestGetTrxData(t *testing.T) {
	coinMarket := CoinexMarket{}

	resp := coinMarket.GetTrxData("BTCBCH", "0")
	fmt.Println(*resp)
}

func TestKline(t *testing.T) {
	coinMarket := CoinexMarket{}

	resp := coinMarket.GetKLine("BTCBCH", "10", "1min")
	fmt.Println(*resp)
}