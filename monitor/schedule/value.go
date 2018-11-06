package main

import (
	"github.com/yqh231/CoinQuantitave/monitor/zb"
	"github.com/yqh231/CoinQuantitave/monitor/decimal"
	"strconv"
)

func topPrice(market string) string{
	var container []string
	kline := zb.Kline(market, "1day", "10")
	for _, data := range *kline.Data{
		container = append(container, strconv.FormatFloat(data[2], 'f', -1, 64))
	}
	return decimal.MaxStr("0", container...)
}

func curPrice(market string) string{
	ticker := zb.Ticker(market)
	return ticker.Ticker.Last
}