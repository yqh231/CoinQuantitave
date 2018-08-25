package trx

import ("net/http"
		"github.com/buger/jsonparser")

type CoinexMarket struct{}

var apiMap = map[string]string{
	"marketList": "https://api.coinex.com/v1/market/list",
	"marketTicker": "https://api.coinex.com/v1/market/ticker",
	"marketTickerAll": "https://api.coinex.com/v1/market/ticker/all"}


func(m *CoinexMarket) GetMarketList() *[]string{
	var marketList []string

	newRequest, _ := http.NewRequest("GET", apiMap["marketList"], nil)
	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
		marketList = append(marketList, string(value))
	}, "data")

	return &marketList
}

func(m *CoinexMarket) GetMarketTicker(marketType string) *map[string]string{
	var marketTicker = make(map[string]string)

	newRequest, _ := http.NewRequest("GET", apiMap["marketTicker"], nil)
	q := newRequest.URL.Query()
	q.Add("market", marketType)
	newRequest.URL.RawQuery = q.Encode()

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ObjectEach(resp, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error{
		marketTicker[string(key)] = string(value)
		return nil
	}, "data", "ticker")

	return &marketTicker
	}


func(m *CoinexMarket) GetMarketTickerAll() *map[string]map[string]string{
	var marketTickerAll = make(map[string]map[string]string)

	newRequest, _ := http.NewRequest("GET", apiMap["marketTickerAll"], nil)

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ObjectEach(resp, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error{
		var valueMap = make(map[string]string)
		jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error{
			valueMap[string(key)] = string(value)
			return nil
		})
		marketTickerAll[string(key)] = valueMap
		return nil
	}, "data", "ticker")

	return &marketTickerAll
	}


func(m *CoinexMarket) 