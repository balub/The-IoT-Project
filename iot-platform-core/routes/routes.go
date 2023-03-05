package routes

import (
	"github.com/balub/The-IoT-Project/controllers"
	"github.com/balub/The-IoT-Project/controllers/client"
	"github.com/balub/The-IoT-Project/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {

	publicRoute := r.Group("/auth")
	protectedRoute := r.Group("/protected")
	hardwareRoute := r.Group("/device")

	protectedRoute.Use(middlewares.JwtAuthMiddleware())

	publicRoute.POST("/register", controllers.HandleRegistration)
	publicRoute.POST("/login", controllers.HandleAuth)

	protectedRoute.POST("/project", controllers.CreateNewProject)
	protectedRoute.GET("/project", controllers.FetchProjects)
	protectedRoute.POST("/device", controllers.CreateNewDevice)
	protectedRoute.GET("/device", controllers.FetchDeviceList)
	protectedRoute.POST("/model", controllers.CreateModel)
	protectedRoute.GET("/model", controllers.FetchModelInfo)
	protectedRoute.POST("/model/edit", controllers.UpdateDataModel)
	protectedRoute.GET("/sse", client.SseHandler)
	protectedRoute.GET("/influx", client.FetchAll)
	hardwareRoute.GET("/model", controllers.FetchSpecificModel)

}
