package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tunamsyar/go-elevator/models"
)

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
}

type AddFloorToUserInput struct {
	FloorID uint `json:"floor_id" binding:"required"`
}

func AddFloorToUser(c *gin.Context) {
	var input AddFloorToUserInput
	var user models.User
	var floor models.Floor

	userID := c.Param("id")

	if err := models.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.First(&floor, input.FloorID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Association("Floors").Append(&floor)

	c.JSON(http.StatusOK, gin.H{"message": "Floor added to user"})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
