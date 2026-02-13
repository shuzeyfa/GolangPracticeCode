package main

import (
	"fmt"
	"time"

	"Task4/controllers"
	"Task4/models"
	"Task4/services"
)

func main() {
	library := services.NewLibrary()

	// add sample members (you already have this)
	library.Member[1] = models.Member{Id: 1, Name: "FirstUser"}
	library.Member[2] = models.Member{Id: 2, Name: "SecondUser"}

	// ── Add a test book so we have something to reserve ──
	library.AddBook(models.Book{
		Id:     100,
		Title:  "Test Book",
		Author: "Test Author",
	})

	// ── Quick test: try to reserve the book with member 1 ──
	err := library.ReserveBook(100, 1)
	if err != nil {
		fmt.Println("Reserve failed:", err)
	} else {
		fmt.Println("Member 1 reserved book 100 successfully")
	}

	// ── Try again immediately with member 2 → should fail ──
	err = library.ReserveBook(100, 2)
	if err != nil {
		fmt.Println("Member 2 reserve failed (good):", err)
	} else {
		fmt.Println("Member 2 reserved — this should NOT happen!")
	}

	// Wait 6 seconds to see auto-cancel
	fmt.Println("Waiting 6 seconds to check auto-release...")
	time.Sleep(6 * time.Second)

	// Try again with member 2 → should succeed now
	err = library.ReserveBook(100, 2)
	if err != nil {
		fmt.Println("Member 2 reserve after timeout failed:", err)
	} else {
		fmt.Println("Member 2 reserved after timeout — success!")
	}

	// Now start the normal menu
	controllers.StartLibrary(library)
}
