package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yqh231/CoinQuantitave/monitor/mongo"
	"net/http"
)

func addMarket(c *gin.Context){
	market := c.PostForm("market") 
	fmt.Println(market)
	if market != ""{
		mongo.MarketInsertOne(market)
	}
	c.JSON(http.StatusOK, ResOk(nil))
}

func listMarket(c *gin.Context){
	var res []string

	markets := mongo.MarketAll()
	for _, market := range markets{
		res = append(res, market.Lookup("market").StringValue())
	}
	c.JSON(http.StatusOK, ResOk(res))
}