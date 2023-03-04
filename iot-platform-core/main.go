package main

import (
	"fmt"

	"github.com/balub/The-IoT-Project/databases"
)

func main() {
    databases.Connect()
    databases.Migrate()
    fmt.Println("Hello, world.")
}