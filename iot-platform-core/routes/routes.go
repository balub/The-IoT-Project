package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	publicRoute := r.Group("/auth")
	// protectedRoute := r.Group("/protected")

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.GET("/login", controllers.HandleAuth)
}
