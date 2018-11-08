package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yqh231/CoinQuantitave/monitor/mongo"
	"net/http"
)

func addMarket(c *gin.Context){
	market := c.DefaultPostForm("market", "")
	mongo.MarketInsertOne(market)

	c.JSON(http.StatusOK, ResOk(nil))
}