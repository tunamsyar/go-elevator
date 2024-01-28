package main

import (
  "github.com/tunamsyar/go-elevator/controllers"
  "github.com/tunamsyar/go-elevator/models"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  models.ConnectDatabase()

  router.POST("/users", controllers.CreateUser)

  router.Run("localhost:8080")
}

