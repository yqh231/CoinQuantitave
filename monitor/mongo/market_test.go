package mongo

import(
	"testing"
	"fmt"
)

/*func TestInsert(t *testing.T){
	cases := []string{"BCH", "ETH", "BTC", "EOS"}
	for _, c := range cases{
		CoinInsertOne(c)
	}
}*/

func TestMarketAll(t *testing.T){
	coins := MarketAll()
	for _, coin := range coins{
		fmt.Println(coin.Lookup("coin").StringValue())
	}
}