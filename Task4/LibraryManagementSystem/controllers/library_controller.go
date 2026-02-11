package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/models"
	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/services"
)

func StartLibrary(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Reserve Book")
		fmt.Println("8. Exit")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice, try again.")
			fmt.Println()
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Book ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Title: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Author: ")
			author, _ := reader.ReadString('\n')

			library.AddBook(models.Book{
				Id:     id,
				Title:  strings.TrimSpace(title),
				Author: strings.TrimSpace(author),
			})

			fmt.Println("Book added successfully")
			fmt.Println()

		case 2:
			fmt.Print("Book Id: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			library.RemoveBook(id)

			fmt.Println("Book Removed Successfully!")
			fmt.Println()

		case 3:
			fmt.Print("Book Id: ")
			bookidStr, _ := reader.ReadString('\n')
			bookid, _ := strconv.Atoi(strings.TrimSpace(bookidStr))

			fmt.Print("Member Id: ")
			memberidStr, _ := reader.ReadString('\n')
			memberid, _ := strconv.Atoi(strings.TrimSpace(memberidStr))

			er := library.BorrowBook(bookid, memberid)

			if er != nil {
				fmt.Println("Error:", er.Error())
				fmt.Println()
			} else {
				fmt.Println("Borrowed successfully!")
				fmt.Println()
			}

		case 4:
			fmt.Print("Book Id: ")
			bookidStr, _ := reader.ReadString('\n')
			bookid, _ := strconv.Atoi(strings.TrimSpace(bookidStr))

			fmt.Print("Member Id: ")
			memberidStr, _ := reader.ReadString('\n')
			memberid, _ := strconv.Atoi(strings.TrimSpace(memberidStr))

			er := library.ReturnBook(bookid, memberid)

			if er != nil {
				fmt.Println("Error:", er.Error())
				fmt.Println()
			} else {
				fmt.Println("Returned successfully!")
				fmt.Println()
			}

		case 5:
			books := library.ListAvailableBooks()
			if len(books) == 0 {
				fmt.Println("No available books.")
			} else {
				for _, b := range books {
					fmt.Printf("Book Id: %d, Title: %s, Author: %s\n", b.Id, b.Title, b.Author)
				}
			}
			fmt.Println()

		case 6:
			fmt.Print("Member Id: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			borrowed := library.ListBorrowedBooks(id)
			if len(borrowed) == 0 {
				fmt.Println("No borrowed books.")
			} else {
				for _, book := range borrowed {
					fmt.Printf("%d %s %s\n", book.Id, book.Title, book.Author)
				}
			}
			fmt.Println()

		case 7: // Reserve Book
			fmt.Print("Book Id: ")
			bookidStr, _ := reader.ReadString('\n')
			bookid, _ := strconv.Atoi(strings.TrimSpace(bookidStr))

			fmt.Print("Member Id: ")
			memberidStr, _ := reader.ReadString('\n')
			memberid, _ := strconv.Atoi(strings.TrimSpace(memberidStr))

			er := library.ReserveBook(bookid, memberid)

			if er != nil {
				fmt.Println("Error:", er.Error())
				fmt.Println()
			} else {
				fmt.Println("Reserved successfully for 5 seconds!")
				fmt.Println()
			}

		case 8:
			fmt.Println("Goodbye")
			fmt.Println()
			return

		default:
			fmt.Println("Invalid option, please choose 1-8.")
			fmt.Println()
		}
	}
}
