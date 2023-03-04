package controllers

import "github.com/gin-gonic/gin"

func HandleRegistration(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "registration",
	})
}