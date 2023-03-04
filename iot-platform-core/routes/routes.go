package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/balub/The-IoT-Project/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	protectedRoute.Use(middlewares.JwtAuthMiddleware())

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.POST("/login", controllers.HandleAuth)

	protectedRoute.POST("/project", controllers.CreateNewProject)
	protectedRoute.GET("/project", controllers.FetchProjects)
}
