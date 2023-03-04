package controllers

import "github.com/gin-gonic/gin"

func HandleAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Signin",
	})
}