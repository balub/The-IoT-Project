package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/databases/models"
	"github.com/balub/The-IoT-Project/utils"
)

type AuthInput struct {
	Username string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HandleAuth(c *gin.Context) {
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Username
	user.Password = input.Password

	u := models.User{}

	dbErr := databases.DB.Model(models.User{}).Where("email = ?", user.Email).Take(&u).Error

	if dbErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbErr.Error()})
		return
	}

	passCheckErr := VerifyPassword(user.Password, u.Password)

	if passCheckErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": passCheckErr.Error()})
		return
	}

	token, tokenErr := utils.GenerateToken(u.ID)

	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
