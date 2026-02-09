package models

const BookAvailable string = "Available"
const BookBorrowed string = "Borrowed"

type Book struct {
	Id     int
	Title  string
	Author string
	Status string
}
