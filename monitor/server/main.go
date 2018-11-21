package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine{
	router := gin.Default()

	market := router.Group("/market")
	{
		market.GET("/bp", breakPoint)
		market.GET("/list", listMarket)
		market.POST("/add", addMarket)
	}

	return router
}

func main(){
	router := setupRouter()
	router.Run()
}