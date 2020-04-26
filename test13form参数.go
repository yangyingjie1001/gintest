package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user/search",func(c *gin.Context)  {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"username":username,
			"address":address,
		})
	})
	r.Run(":8080")
}