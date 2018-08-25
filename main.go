package main 

import (
	"quan/trx"
)

func main() {
	var market = trx.CoinexMarket{}

	market.GetMarketList()
}