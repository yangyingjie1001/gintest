package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON",func(c *gin.Context){
		// 方式一：自己拼接json
		c.JSON(http.StatusOK,gin.H{"message":"hello world"})
	})
	r.GET("/moreJSON",func(c *gin.Context){
		var msg struct {
			Name string `json:"user"`
			Message string
			Age int
		}
		msg.Name = "小王子"
		msg.Message = "hello world"
		msg.Age = 18
		c.JSON(http.StatusOK ,msg)
	})
	r.Run(":8080")
}