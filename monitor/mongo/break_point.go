package mongo

/*
{
	"_id": ObjectId(),
	"create_time": int64,
	"market": "bch_usdt"
	"cur_price": string,
	"top_price": string,
	"break_point": bool,
}
*/

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
	l "github.com/yqh231/CoinQuantitave/monitor/log"
)

const breakPoint = "break_point"

func BpInsertOne(market, cur, top string, bp bool){
	c := Mgo.Use(breakPoint)
	doc := []map[string]interface{}{
		{"cur_price": cur},
		{"top_price": top},
		{"break_point": bp},
		{"market": market},
		{"create_time": time.Now().Unix()},
	}
	_, err := c.InsertOne(context.Background(), MakeDoc(doc))
	if err != nil{
		l.Logger.Println(err.Error())
	}
}

func FilterBp(filter *bson.Document) (r []*bson.Document){
	c := Mgo.Use(breakPoint)

	cur, _ := c.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	for cur.Next(context.Background()){
		doc := bson.NewDocument()
		cur.Decode(doc)
		r = append(r, doc)
	}
	return
}