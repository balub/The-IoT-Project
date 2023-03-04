package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
)

type RegisterInput struct {
	Username string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleRegistration(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{" Input error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Username
	u.Password = input.Password

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{" Password error": errHash.Error()})
		return
	}

	u.Password = string(hashedPassword)

	u.ID = cuid.New()

	errCase := databases.DB.Create(&u).Error

	if errCase != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Hash error": errCase.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
