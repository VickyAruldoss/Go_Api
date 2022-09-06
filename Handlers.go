package main

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers ... Get all users
// @Summary Get all users
// @Description get all users
// @Tags Users
// @Success 200  {object} model.User
// @Failure 404 {object} object
// @Router / [get]
func GetUsers(c *gin.Context) {

	user, err := model.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Users": user})
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	model.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"message": "savedSuccessFully"})
	return
}

func GetUserById(c *gin.Context) {
	//id := c.Params.ByName("id")

}
