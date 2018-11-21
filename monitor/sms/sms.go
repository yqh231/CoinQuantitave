package sms

import (
	"fmt"

	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	l "github.com/yqh231/CoinQuantitave/monitor/log"
)

func Send(coin string) {
	clnt := ypclnt.New("10621b0024408c05f6c57a4007872a7c")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = "18689468297"
	param[ypclnt.TEXT] = fmt.Sprintf("【杨启航】您喜欢的商品%s，价格发生了变化，请知悉。", coin)
	r := clnt.Sms().SingleSend(param)
	if r.Code != ypclnt.SUCC{
		l.Logger.Println(r.Msg)
	}
}

