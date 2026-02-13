package main

import (
	"TaskManagement/router"
	"fmt"
)

func main() {
	r := router.Router()
	fmt.Println("Task Management API is running...")

	r.Run(":8080")
}
