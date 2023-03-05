package main

import (
	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// handle database connections
	databases.Connect()
	databases.Migrate()

	// setup http server
	router := gin.Default()

	// router.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
	// 	c.Next()
	// })

	router.Use(cors.Default())

	routes.SetUpRouter(router)

	router.Run("localhost:9090")
}
