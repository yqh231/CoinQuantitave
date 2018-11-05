package mongo

import (
	"testing"
	"context"
	"github.com/okcoin-okex/open-api-v3-sdk/okex-go-sdk-api"
	"fmt"
)

func TestConnect(t *testing.T) {
	collenct_name := "test"
	New()
	c := Mgo.Use(collenct_name)
	c.InsertOne(context.Background(), map[string]string{"hello": "world"})
}


func TestOKExServerTime(t *testing.T) {
	serverTime, err := NewOKExClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("OKEx's server time: ", serverTime)
}

func NewOKExClient() *okex.Client {
	var config okex.Config
	config.Endpoint = "https://www.okex.cn/"
	config.ApiKey = "" 
	config.SecretKey = "" 
	config.Passphrase = "" 
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = okex.ENGLISH

	client := okex.NewClient(config)
	return client
}