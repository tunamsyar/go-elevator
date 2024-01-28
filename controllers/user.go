package controllers

import (
  "net/http"

  "github.com/tunamsyar/go-elevator/models"
  "github.com/gin-gonic/gin"
)

type CreateUserInput struct {
  Name string `json:"name" binding:"required"`
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
