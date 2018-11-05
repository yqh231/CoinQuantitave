package zb

import (
	"testing"
	"fmt"
	"time"
)

func TestKline(t *testing.T){
	data := Kline("bch_usdt", "1day", "10")
	for _, data := range *data.Data{
		fmt.Println(time.Unix(int64(data[0]/1e3), 0).Format("2006-01-02 15:04:05"))
	}
}

func TestTicker(t *testing.T){
	r := Ticker("btc_usdt")
	fmt.Println(*r)
}
