package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"fullname": "Vicky Aruldoss",
}
var us = []string{"Vicky", "Aruldoss", "John", "Smith"}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", GetUsers)
	r.GET("/getallusers", GetAllUsers)
	//r.POST("/user",CreateUser)
	//r.DELETE("/user",DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"Users": us,
	})
}

func GetUsers(c *gin.Context) {
	name := c.Query("fullname")
	fullname, ok := users[name]
	if !ok {
		c.JSON(404, gin.H{
			name: "",
		})
		return
	}
	c.JSON(200, gin.H{
		name: fullname,
	})
}
