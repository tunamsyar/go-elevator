package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tunamsyar/go-elevator/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Mercu 2")
	})

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.FindUsers)
	router.GET("/users/:id", controllers.FindUser)
	router.Run("localhost:8080")
}
