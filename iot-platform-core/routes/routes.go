package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine){
	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")

	publicRoute.GET("/",controllers.HandleAuth)
	protectedRoute.GET("/",controllers.HandleRegistration)

	protectedRoute.POST("/createproject", controllers.CreateNewProject)
	protectedRoute.GET("/getprojects",controllers.FetchProjects)

	
}