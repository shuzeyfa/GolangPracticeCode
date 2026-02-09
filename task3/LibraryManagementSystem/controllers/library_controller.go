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
		fmt.Println("6. List Borrowed Book")
		fmt.Println("7. Exit")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

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

			fmt.Println("Book Removed Succefully!")
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
			for _, b := range books {
				fmt.Println("Book Id: ", b.Id, "Book Title: ", b.Title, "Book Author: ", b.Author)
			}
			fmt.Println()
		case 6:
			fmt.Print("Member Id: ")

			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			borrowed := library.ListBorrowedBooks(id)

			for _, book := range borrowed {
				fmt.Println(book.Id, book.Title, book.Author)
			}
			fmt.Println()

		case 7:
			fmt.Println("Goodbye")
			fmt.Println()
			return
		}
	}
}
