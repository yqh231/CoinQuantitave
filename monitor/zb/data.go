package zb

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const EndPoint = "http://api.zb.cn/data/v1/"

type data [][]float64

type kline struct{
	Data *data
	MoneyType string
	Symbol string
}

type ticker struct{
	Ticker struct{
		Vol string
		Last string
		Sell string
		Buy string
		High string
		Low string
	}
	Date string
}

var client = &http.Client{}

func DoRequest(req *http.Request) []byte{
	resp, err := client.Do(req) 
	if err != nil{

	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}


func Kline(market, type_, size string) *kline{
	var kline = new(kline)
	newRequest, _ := http.NewRequest("GET", EndPoint+"kline", nil)

	q := newRequest.URL.Query()
	q.Add("market", market)
	q.Add("type", type_)
	q.Add("size", size)
	newRequest.URL.RawQuery = q.Encode()

	resp := DoRequest(newRequest)

	json.Unmarshal(resp, kline)
	return kline
}

func Ticker(market string) *ticker{
	var ticker = new(ticker)
	newRequest, _ := http.NewRequest("GET", EndPoint+"ticker", nil)

	q := newRequest.URL.Query()
	q.Add("market", market)
	newRequest.URL.RawQuery = q.Encode()

	resp := DoRequest(newRequest)

	json.Unmarshal(resp, ticker)
	return ticker 
}