package library

import (
	"time"
)

type Book struct {
	Title         string
	Author        string
	NumberOfPages int
	IsFinished    bool
	AddTime       time.Time
	FinishTime    *time.Time
}

func NewBook(title string, author string, numberOfPages int) (Book, error) {
	if title == "" {
		return Book{}, ErrInvalidBookTitle
	}
	if author == "" {
		return Book{}, ErrInvalidBookAuthor
	}
	if numberOfPages <= 0 {
		return Book{}, ErrInvalidBookNumberOfPages
	}

	return Book{
		Title:         title,
		Author:        author,
		NumberOfPages: numberOfPages,
		IsFinished:    false,
		AddTime:       time.Now(),
		FinishTime:    nil,
	}, nil
}
