package services

import (
	"errors"

	"github.com/shuzeyfa/GolangPracticeCode/task3/LibraryManagementSystem/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookId int)
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
}

type Library struct {
	Books  map[int]models.Book
	Member map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:  make(map[int]models.Book),
		Member: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	book.Status = models.BookAvailable
	l.Books[book.Id] = book
}

func (l *Library) RemoveBook(bookId int) {

	_, exists := l.Books[bookId]
	if exists {
		delete(l.Books, bookId)
	}
}

func (l *Library) ListAvailableBooks() []models.Book {

	available := []models.Book{}

	for _, book := range l.Books {
		if book.Status == models.BookAvailable {
			available = append(available, book)
		}
	}

	return available
}

func (l *Library) ListBorrowedBooks(memberId int) []models.Book {

	member, memberExist := l.Member[memberId]
	if !memberExist {
		return []models.Book{}
	}

	return member.BorrowedBooks
}

func (l *Library) BorrowBook(bookId int, memberId int) error {

	book, bookExist := l.Books[bookId]
	if !bookExist {
		return errors.New("Book Not Found!")
	}

	member, memberExist := l.Member[memberId]
	if !memberExist {
		return errors.New("Member Does Not Exist!")
	}

	if book.Status == models.BookBorrowed {
		return errors.New("Book already borrowed!")
	}

	book.Status = models.BookBorrowed
	l.Books[bookId] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Member[memberId] = member

	return nil

}

func (l *Library) ReturnBook(bookId int, memberId int) error {

	book, bookExist := l.Books[bookId]
	if !bookExist {
		return errors.New("Book not found!")
	}

	member, memberExist := l.Member[memberId]
	if !memberExist {
		return errors.New("Member not found!")
	}

	updatedBook := []models.Book{}

	check := false

	for _, book := range member.BorrowedBooks {
		if book.Id == bookId {
			check = true
			continue
		}
		updatedBook = append(updatedBook, book)
	}

	if !check {
		return errors.New("This member doesnot  borrow this book!")
	}

	member.BorrowedBooks = updatedBook
	l.Member[memberId] = member

	book.Status = models.BookAvailable
	l.Books[bookId] = book

	return nil

}
