package trx

import (
	"fmt"
	"testing"
)

func TestGetHuobiMarketDetail(t *testing.T){
	huobiApi := HuobiMarket{}

	model:=huobiApi.GetHuobiMarketDetail("veteth")
	fmt.Println((*model).Tick.Amount)
}