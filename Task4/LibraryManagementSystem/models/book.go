package models

import "time"

const BookAvailable string = "Available"
const BookBorrowed string = "Borrowed"
const BookReserved string = "Reserved"

type Book struct {
	Id            int
	Title         string
	Author        string
	Status        string
	ReservedBy    int
	ReservedUntil time.Time
}
