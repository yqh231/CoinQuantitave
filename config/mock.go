package config

var TrxMarketType = []string{"VETETH"}

var MockData = map[string]map[string]string{
	"VETETH": map[string]string{"huobiEth": "5", "huobiVet": "800", "huobiRate": "0.001", "coinexRate": "0",
								"coinexEth": "5", "coinexVet": "800", "tradePercent": "0.05"},
}