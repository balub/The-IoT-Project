package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	protectedRoute.POST("/project", controllers.CreateNewProject)
	protectedRoute.GET("/project", controllers.FetchProjects)

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.POST("/login", controllers.HandleAuth)
}
