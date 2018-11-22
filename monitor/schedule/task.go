package main

import(
	"time"

	"github.com/yqh231/CoinQuantitave/monitor/decimal"
	"github.com/yqh231/CoinQuantitave/monitor/mongo"
	"github.com/yqh231/CoinQuantitave/monitor/sms"
)


func taskBreakPoint(){
	markets := mongo.MarketAll()
	var flag, low_flag bool
	for _, market:= range markets{
		top, low := Price(market.Lookup("market").StringValue())
		cur := curPrice(market.Lookup("market").StringValue())
		if decimal.CompareStr(cur, top) == 1{
			flag = true	
			sms.Send(market.Lookup("market").StringValue())
		}

		if decimal.CompareStr(cur, low) == -1{
			low_flag = true
			sms.Send(market.Lookup("market").StringValue())
		}
		mongo.BpInsertOne(market.Lookup("market").StringValue(), cur, top, flag)
		mongo.LowInsertOne(market.Lookup("market").StringValue(), cur, low, low_flag)
		time.Sleep(5000 * time.Millisecond)
	}
}