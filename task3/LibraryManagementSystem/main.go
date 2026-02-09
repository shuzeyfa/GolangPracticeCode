package main

import (
	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/controllers"
	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/models"
	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/services"
)

func main() {
	library := services.NewLibrary()

	// add sample members
	library.Member[1] = models.Member{Id: 1, Name: "Alice"}
	library.Member[2] = models.Member{Id: 2, Name: "Bob"}

	controllers.StartLibrary(library)
}
