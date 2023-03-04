package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			token := c.GetHeader("Authorization")
			if token == "" {
					c.Set("userId", nil)
			} else {
					// Perform authentication logic here
					// ...
					// Set auth value based on authentication success or failure
					c.Set("userId", 1)
			}
			c.Next()
	}
}

func SetUpRouter(r *gin.Engine){
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	publicRoute.GET("/",controllers.HandleAuth)
	protectedRoute.GET("/",controllers.HandleRegistration)


	protectedRoute.Use(AuthMiddleware())
	protectedRoute.POST("/createProject", controllers.CreateNewProject)

	
}