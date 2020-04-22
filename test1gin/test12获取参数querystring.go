package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
	// get 请求
func main() {
	r := gin.Default()
	r .GET("user/search",func(c *gin.Context){
		// username := c.DefaultQuery("username","小王子")
		username := c.Query("username")
		address := c.Query("address")
		c.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"username":username,
			"address":address,
		})
	})
	r.Run(":8080")
}