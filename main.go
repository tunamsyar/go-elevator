package main

import (
	"github.com/tunamsyar/go-elevator/models"
	"github.com/tunamsyar/go-elevator/router"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	ginEngine := gin.Default()
	router.SetupRoutes(ginEngine)

}
