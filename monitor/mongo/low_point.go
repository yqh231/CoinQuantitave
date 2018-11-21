package mongo

/*
{
	"_id": ObjectId(),
	"create_time": int64,
	"market": "bch_usdt"
	"cur_price": string,
	"low_price": string,
	"break_point": bool,
}
*/

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
	l "github.com/yqh231/CoinQuantitave/monitor/log"
)

const lowPoint = "low_point"

func LowInsertOne(market, cur, low string, bp bool){
	c := Mgo.Use(lowPoint)
	doc := []map[string]interface{}{
		{"cur_price": cur},
		{"low_price": low},
		{"break_point": bp},
		{"market": market},
		{"create_time": time.Now().Unix()},
	}
	_, err := c.InsertOne(context.Background(), MakeDoc(doc))
	if err != nil{
		l.Logger.Println(err.Error())
	}
}

func Filterlow(filter *bson.Document) (r []*bson.Document){
	c := Mgo.Use(lowPoint)

	cur, _ := c.Find(context.Background(), filter)
	defer func(){
		if cur != nil{
			cur.Close(context.Background())
		}
	}()
	for cur.Next(context.Background()){
		doc := bson.NewDocument()
		cur.Decode(doc)
		r = append(r, doc)
	}
	return
}