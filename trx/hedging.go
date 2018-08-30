package trx

import (
	"fmt"
	conf "quan/config"
	"strings"
	//mgo "quan/trace_ctl"
	"github.com/shopspring/decimal"
)


type Hedging struct{
}

func stringToDecimal(str string) decimal.Decimal{
	decimal_, _ := decimal.NewFromString(str)
	return decimal_
}

func getServiceCharge(market string, prices ...decimal.Decimal) decimal.Decimal{
	total, _ := decimal.NewFromString("0")
	for _, price := range prices{
		total = total.Add(price)
	}
	return total.Mul(stringToDecimal(conf.MockData[market]["huobiRate"]))
}

func(h *Hedging) Hedge() {
	m := CoinexMarket{}
	huobi := HuobiMarket{}

	for _, market := range conf.TrxMarketType{
		coinEx := m.GetMarketTicker(market)
		huoBi := huobi.GetHuobiMarketDetail(strings.ToLower(market))
		coinExPrice, _ := decimal.NewFromString((*coinEx)["last"])
		huoBiPrice := decimal.NewFromFloat((*huoBi).Tick.Close)
		if coinExPrice.Cmp(huoBiPrice) == -1{
			diffValue := huoBiPrice.Sub(coinExPrice)
			fmt.Println(diffValue)
			if diffValue.GreaterThan(getServiceCharge(market, coinExPrice, huoBiPrice)) {
				//trade
			}
		}else if coinExPrice.Cmp(huoBiPrice) == 1{
			diffValue := coinExPrice.Sub(huoBiPrice)
			fmt.Println(diffValue)
			if diffValue.GreaterThan(getServiceCharge(market, coinExPrice, huoBiPrice)){
				//trade
			}
		}
	}
}
