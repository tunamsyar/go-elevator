package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunamsyar/go-elevator/models"
)

type CreateFloorInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateFloor(c *gin.Context) {
	var input CreateFloorInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	floor := models.Floor{Name: input.Name}
	models.DB.Create(&floor)

	c.JSON(http.StatusOK, gin.H{"data": floor})
}

func FindFloors(c *gin.Context) {
	var floors []models.Floor

	models.DB.Find(&floors)

	c.JSON(http.StatusOK, gin.H{"data": floors})
}

func FindFloor(c *gin.Context) {
	var floor models.Floor

	if err := models.DB.Where("id = ?", c.Param("id")).First(&floor).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": floor})
}
