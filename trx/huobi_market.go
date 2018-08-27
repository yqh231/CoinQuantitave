package trx

import (
	"net/http"
	"encoding/json"
)

type HuobiMarket struct{}

var huoBiApi = map[string]string{
	"marketDetail": "https://api.huobipro.com/market/detail"}


func(h *HuobiMarket) GetHuobiMarketDetail(marketType string) *MarketDetailReturn{
	marketDetailReturn := MarketDetailReturn{}

	newRequest, _ := http.NewRequest("GET", huoBiApi["marketDetail"], nil)

	q := newRequest.URL.Query()
	q.Add("symbol", marketType)
	newRequest.URL.RawQuery = q.Encode()

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()
	json.Unmarshal(resp, &marketDetailReturn)

	return &marketDetailReturn
}