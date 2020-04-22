package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someXML",func(c *gin.Context){
		c.XML(http.StatusOK,gin.H{"message":"Hello world!"})
	})
	r.GET("/moreXML",func (c *gin.Context){
		// 方法二 使用结构体
		type MessageRecord struct {
			Name string
			Message string
			Age int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	}) 
	r.Run(":8080")
}
