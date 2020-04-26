package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Login struct {
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
func main() {
	r := gin.Default()
	// 绑定JSON的示列({"user":"q1mi","password":"123456"})
	r.POST("/loginJSON",func(c *gin.Context){
		var login Login
		if err := c.ShouldBind(&login);err == nil{
			fmt.Printf("login info:%#v\n",login)
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password":login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}
	})
	r.POST("loginForm",func(c *gin.Context){
		var login Login
		err := c.ShouldBind(&login)
		if err ==nil {
			fmt.Printf("login info:%#v\n",login)
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password":login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		}
	})
	r.GET("loginForm",func(c *gin.Context){
		var login Login
		err := c.ShouldBind(&login)
		if err ==nil {
			fmt.Printf("login info:%#v\n",login)
			c.JSON(http.StatusOK,gin.H{
				"user": login.User,
				"password":login.Password,
			})
		}else{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		}
	})
	r.Run(":8080")
}