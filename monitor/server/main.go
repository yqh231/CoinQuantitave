package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	market := router.Group("/market")
	{
		market.GET("/bp", breakPoint)
		market.PUT("/bp", addMarket)
	}

	templates := router.Group("/")
	{
		templates.GET("", base)
	}
	return router
}

func main(){
	router := setupRouter()
	router.Run()
}