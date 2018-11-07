package mongo

import(
	"testing"
	"fmt"
)

func TestInsert(t *testing.T){
	cases := []string{"bch_usdt", "eth_usdt", "btc_usdt", "eos_usdt"}
	for _, c := range cases{
		MarketInsertOne(c)
	}
}

func TestMarketAll(t *testing.T){
	coins := MarketAll()
	for _, coin := range coins{
		fmt.Println(coin.Lookup("market").StringValue())
	}
}