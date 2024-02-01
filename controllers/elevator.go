package controllers

import (
	"net/http"
  "fmt"

	"github.com/gin-gonic/gin"
	"github.com/tunamsyar/go-elevator/models"
)

type FloorSelectionInput struct {
	FloorID uint `json:"floor_id" binding:"required"`
}

func FloorSelection(c *gin.Context) {
	var input FloorSelectionInput
	var floor models.Floor
	var user models.User
  
  userID := c.Param("id")
	
  if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := models.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.First(&floor, input.FloorID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	associated, err := isUserAssociatedWithFloor(user, floor)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !associated {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User is not authorized with floor"})
		return
	}

	elevator := models.NewElevator()
	elevator.AddFloor(int(floor.ID))

	c.JSON(http.StatusOK, gin.H{"message": "Floor added to elevator queue"})
}

func isUserAssociatedWithFloor(user models.User, floor models.Floor) (bool, error) {
	var count int64
	result := models.DB.Model(&user).Where("id = ? AND ? IN (SELECT floor_id FROM user_floors WHERE user_id = ?)", user.ID, floor.ID, user.ID).Count(&count)

  fmt.Println("User: ", user.ID)
  fmt.Println("Floor: ", floor.ID)
  fmt.Println("COUNT: ", count)
	
  if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}
