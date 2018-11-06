package mongo

import (
	"testing"
	"context"
)

func TestConnect(t *testing.T) {
	collenct_name := "test"
	New()
	c := Mgo.Use(collenct_name)
	c.InsertOne(context.Background(), map[string]string{"hello": "world"})
}
