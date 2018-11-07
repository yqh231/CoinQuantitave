package mongo

import(
	"testing"
)

func TestBpInsert(t *testing.T){
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "200", true)
	BpInsertOne("bch_usdt", "100", "0", false)
	BpInsertOne("bch_usdt", "100", "200", false)
}

func TestFilterBp(t *testing.T){
	FilterBp(nil)
	FilterBp(MakeDoc([]map[string]interface{}{{"market": "bch_usdt"}}))
}