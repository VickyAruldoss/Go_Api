package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title User API documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /user
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", GetUsers)
	// r.GET("/getallusers", GetAllUsers)
	r.POST("/user", CreateUser)
	//r.DELETE("/user",DeleteUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// r := route.SetupRouter()
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run(":" + os.Getenv("SERVER_PORT"))
}
