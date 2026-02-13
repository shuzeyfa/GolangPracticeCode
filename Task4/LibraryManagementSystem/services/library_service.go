package services

import (
	"errors"
	"sync"
	"time"

	"Task4/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookId int)
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
	ReserveBook(bookId, memberId int) error
}

type Library struct {
	Books        map[int]models.Book
	Member       map[int]models.Member
	mu           sync.Mutex
	reservations chan ReservationRequest
}

type ReservationRequest struct {
	BookId   int
	MemberId int
	Reply    chan error
}

func NewLibrary() *Library {
	lib := &Library{
		Books:        make(map[int]models.Book),
		Member:       make(map[int]models.Member),
		reservations: make(chan ReservationRequest, 50), // buffered = better
	}

	// Start worker
	go lib.reservationWorker()

	return lib
}

func (l *Library) reservationWorker() {
	for req := range l.reservations {
		err := l.tryReserveBook(req.BookId, req.MemberId)
		req.Reply <- err
	}
}

func (l *Library) tryReserveBook(bookID, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status != models.BookAvailable {
		return errors.New("book not available")
	}

	book.Status = models.BookReserved
	book.ReservedBy = memberID
	book.ReservedUntil = time.Now().Add(5 * time.Second)
	l.Books[bookID] = book

	// Auto-release timer
	go func(bID, mID int) {
		time.Sleep(5 * time.Second)
		l.mu.Lock()
		defer l.mu.Unlock()

		b, ok := l.Books[bID]
		if ok && b.Status == models.BookReserved && b.ReservedBy == mID {
			b.Status = models.BookAvailable
			b.ReservedBy = 0
			l.Books[bID] = b
		}
	}(bookID, memberID)

	return nil
}

func (l *Library) ReserveBook(bookId, memberId int) error {
	l.mu.Lock()
	if _, ok := l.Member[memberId]; !ok {
		l.mu.Unlock()
		return errors.New("member not found")
	}
	l.mu.Unlock()

	reply := make(chan error, 1)
	l.reservations <- ReservationRequest{
		BookId:   bookId,
		MemberId: memberId,
		Reply:    reply,
	}

	return <-reply
}

func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	book.Status = models.BookAvailable
	l.Books[book.Id] = book
}

func (l *Library) RemoveBook(bookId int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.Books, bookId)
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	var available []models.Book
	for _, book := range l.Books {
		if book.Status == models.BookAvailable {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberId int) []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, ok := l.Member[memberId]
	if !ok {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

func (l *Library) BorrowBook(bookId int, memberId int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookId]
	if !exists {
		return errors.New("Book Not Found!")
	}

	member, exists := l.Member[memberId]
	if !exists {
		return errors.New("Member Does Not Exist!")
	}

	if book.Status == models.BookBorrowed {
		return errors.New("Book already borrowed!")
	}

	if book.Status == models.BookReserved {
		if book.ReservedBy != memberId {
			return errors.New("Book reserved by another member!")
		}
		book.ReservedBy = 0
	}

	book.Status = models.BookBorrowed
	l.Books[bookId] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Member[memberId] = member

	return nil
}

func (l *Library) ReturnBook(bookId int, memberId int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.Books[bookId]
	if !exists {
		return errors.New("Book not found!")
	}

	member, exists := l.Member[memberId]
	if !exists {
		return errors.New("Member not found!")
	}

	var updated []models.Book
	found := false
	for _, b := range member.BorrowedBooks {
		if b.Id == bookId {
			found = true
			continue
		}
		updated = append(updated, b)
	}

	if !found {
		return errors.New("This member does not borrow this book!")
	}

	member.BorrowedBooks = updated
	l.Member[memberId] = member

	book.Status = models.BookAvailable
	l.Books[bookId] = book

	return nil
}
