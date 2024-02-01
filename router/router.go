package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tunamsyar/go-elevator/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Mercu 2")
	})

	router.POST("/floors", controllers.CreateFloor)
	router.GET("/floors", controllers.FindFloors)
	router.GET("/floors/:id", controllers.FindFloor)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.FindUsers)
	router.GET("/users/:id", controllers.FindUser)
	router.POST("/users/:id/add_floor", controllers.AddFloorToUser)
  router.POST("/users/:id/floor_selection", controllers.FloorSelection)
	router.Run("localhost:8080")
}
