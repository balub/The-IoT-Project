package main

import (
	"fmt"

	"github.com/balub/The-IoT-Project/databases"
	"github.com/balub/The-IoT-Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // handle database connections
    databases.Connect()
    databases.Migrate()

    // setup http server
    router := gin.Default()
    routes.SetUpRouter(router)

    router.Run("localhost:9090")
    

    fmt.Println("Hello, world.")
}