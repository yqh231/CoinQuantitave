package mongo
/*
db design
{
	"_id": ObjectId(),
	"market": "btc_usdt"
}

*/
import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
)

const market_ = "market"

func MarketInsertOne(market string){
	c := Mgo.Use(market_)
	_, err := c.InsertOne(context.Background(), MakeDoc([]map[string]interface{}{{"market": market}}))
	if err != nil{
		//to do
	}
}

func MarketAll() (r []*bson.Document){
	c := Mgo.Use(market_)

	cur, _ := c.Find(context.Background(), nil)
	for cur.Next(context.Background()){
		doc := bson.NewDocument()
		cur.Decode(doc)
		r = append(r, doc)
	}
	return
}

