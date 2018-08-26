package trx

import ("net/http"
		"github.com/buger/jsonparser")

type CoinexMarket struct{}

var apiMap = map[string]string{
	"marketList": "https://api.coinex.com/v1/market/list",
	"marketTicker": "https://api.coinex.com/v1/market/ticker",
	"marketTickerAll": "https://api.coinex.com/v1/market/ticker/all",
	"marketDepth": "https://api.coinex.com/v1/market/depth",
	"trxData": "https://api.coinex.com/v1/market/deals",
	"kLine": "https://api.coinex.com/v1/market/kline"}


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


func(m *CoinexMarket) GetMarketDepth(marketType string, merge string, limit string) *map[string][][]string {
	var marketDepth = make(map[string][][]string)

	newRequest, _ := http.NewRequest("GET", apiMap["marketDepth"], nil)

	q := newRequest.URL.Query()
	q.Add("market", marketType)
	q.Add("merge", merge)
	q.Add("limit", limit)
	newRequest.URL.RawQuery = q.Encode()

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ObjectEach(resp, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error{
		var valueSlice [][]string
		if string(key) == "asks" || string(key) == "bids"{
			jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				var tempContainer []string
				jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
					tempContainer = append(tempContainer, string(value))
				})
				valueSlice = append(valueSlice, tempContainer)
			})
		}
		if valueSlice != nil{
			marketDepth[string(key)] = valueSlice
		}
		return nil
	}, "data")
	return &marketDepth
}

func(m *CoinexMarket) GetTrxData(marketType string, last_id string) *[]map[string]string {
	var trxData []map[string]string

	newRequest, _ := http.NewRequest("GET", apiMap["trxData"], nil)

	q := newRequest.URL.Query()
	q.Add("market", marketType)
	q.Add("last_id", last_id)
	newRequest.URL.RawQuery = q.Encode()

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
		var valueMap = make(map[string]string)

		jsonparser.ObjectEach(value ,func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error{
			valueMap[string(key)] = string(value)
			return nil
		})
		trxData = append(trxData, valueMap)

	}, "data")
	return &trxData
}

func(m *CoinexMarket) GetKLine(marketType, limit, kType string) *[][]string{
	var kLine [][]string

	newRequest, _ := http.NewRequest("GET", apiMap["kLine"], nil)

	q := newRequest.URL.Query()
	q.Add("market", marketType)
	q.Add("limit", limit)
	q.Add("type", kType)
	newRequest.URL.RawQuery = q.Encode()

	c := CoinexRequest{req: newRequest}
	c.SetHeader()
	resp := c.DoRequest()

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error){
		var valueSlice []string
	jsonparser.ArrayEach(value, func(childValue []byte, dataType jsonparser.ValueType, offset int, err error){
		valueSlice = append(valueSlice, string(childValue))
	})
		kLine = append(kLine, valueSlice)
	}, "data")
	return &kLine
}