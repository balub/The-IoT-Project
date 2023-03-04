package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/balub/The-IoT-Project/controllers/client"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	protectedRoute.POST("/createproject", controllers.CreateNewProject)
	protectedRoute.GET("/getprojects", controllers.FetchProjects)

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.POST("/login", controllers.HandleAuth)

	protectedRoute.GET("/sse", client.SseHandler)
}
