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
	c.JSON(http.StatusOK, gin.H{"User": user})
}

// CreateUser ... Create User
// func CreateUser(c *gin.Context) {
// 	var user model.User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := model.CreateUser(&user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "success"})
// }

// GetUserByID ... Get the user by id
// func GetUserByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	userID, err := model.StringToBinaryUUID(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var user model.User
// 	err = model.GetUserByID(&user, userID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }
