package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	market := router.Group("/market")
	{
		market.GET("/bp", breakPoint)
	}
	router.Run()
}