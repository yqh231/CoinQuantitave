package mongo
/*
db design
{
	"_id": ObjectId(),
	"coin": "BTC"
}

*/
import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
)

const market = "market"

func MarketInsertOne(market string){
	c := Mgo.Use(market)
	_, err := c.InsertOne(context.Background(), makeDoc([]map[string]interface{}{{"market": market}}))
	if err != nil{
		//to do
	}
}

func MarketAll() (r []*bson.Document){
	c := Mgo.Use(market)

	cur, _ := c.Find(context.Background(), nil)
	defer cur.Close(context.Background())
	for cur.Next(context.Background()){
		doc := bson.NewDocument()
		cur.Decode(doc)
		r = append(r, doc)
	}
	return
}

