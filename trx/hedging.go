package trx

import (
	"fmt"
	conf "quan/config"
	//mgo "quan/trace_ctl"
	huobi "github.com/huobiapi/REST-GO-demo/services"
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

	for _, market := range conf.TrxMarketType{
		coinEx := m.GetMarketTicker(market)
		huoBi := huobi.GetTradeDetail(market)

		coinExPrice, _ := decimal.NewFromString((*coinEx)["latest"])
		huoBiPrice := decimal.NewFromFloat(huoBi.Tick.Data[0].Price)
		
		if coinExPrice.Cmp(huoBiPrice) == -1{
			diffValue := huoBiPrice.Sub(coinExPrice)
			if diffValue.GreaterThan(getServiceCharge(market, coinExPrice, huoBiPrice)) {
				//trade
			}
		}else if coinExPrice.Cmp(huoBiPrice) == 1{
			diffValue := coinExPrice.Sub(huoBiPrice)
			if diffValue.GreaterThan(getServiceCharge(market, coinExPrice, huoBiPrice)){
				//trade
			}
		}
	}
}
