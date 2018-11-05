package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
)

const DB = "Monitor"

type MongoManager struct {
	DB      string
	Session *mongo.Client
}

var Mgo *MongoManager
var ColMap = make(map[string]*mongo.Collection)

func New() {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		panic("start mongo fail")
	}
	client.Connect(context.TODO())
	Mgo = &MongoManager{DB, client}
}

func (mgo *MongoManager) Use(collection string) *mongo.Collection{
	var c *mongo.Collection
	var ok bool

	if c, ok = ColMap[collection]; !ok{
		c = mgo.Session.Database(DB).Collection(collection)
		ColMap[collection] = c
		return c
	}
	return c
}
