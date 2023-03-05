package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/balub/The-IoT-Project/controllers/client"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	publicRoute.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	protectedRoute.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// protectedRoute.Use(middlewares.JwtAuthMiddleware())

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.POST("/login", controllers.HandleAuth)

	protectedRoute.POST("/project", controllers.CreateNewProject)
	protectedRoute.GET("/project", controllers.FetchProjects)
	protectedRoute.POST("/device", controllers.CreateNewDevice)
	protectedRoute.GET("/device", controllers.FetchDeviceList)

	// client
	protectedRoute.GET("/sse", client.SseHandler)
	protectedRoute.GET("/influx", client.FetchAll)

}
