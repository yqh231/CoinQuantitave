package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func base(c *gin.Context){
	c.HTML(http.StatusOK, "base.html", nil)
}