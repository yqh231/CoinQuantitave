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