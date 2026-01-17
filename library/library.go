package library

import (
	"sync"
	"time"
)

type Library struct {
	books map[string]Book
	mtx   sync.RWMutex
}

// AddBook adds the provided book to the library.
//
// It returns ErrBookAlreadyExists if the book already exists.
func (l *Library) AddBook(book Book) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.books[book.Title]; ok {
		return ErrBookAlreadyExists
	}

	l.books[book.Title] = book
	return nil
}

// FinishBook marks the book with given title as finished.
//
// It returns ErrBookNotFound if the book does not exist.
// It returns ErrBookIsAlreadyFinished if the book is already finished.
func (l *Library) FinishBook(title string) (Book, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	book, ok := l.books[title]
	if !ok {
		return Book{}, ErrBookNotFound
	}
	if book.IsFinished {
		return Book{}, ErrBookIsAlreadyFinished
	}

	now := time.Now()
	book.IsFinished = true
	book.FinishTime = &now
	l.books[title] = book

	return book, nil
}

// GetBook returns a book with given title if the book exists.
//
// It returns ErrBookNotFound if the book does not exist.
func (l *Library) GetBook(title string) (Book, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	book, ok := l.books[title]

	if !ok {
		return Book{}, ErrBookNotFound
	}

	return book, nil
}

// GetBooks returns books filtered by author and isFinished status.
//
// If author is an empty string, no filtering by author is applied.
// If isFinished is nil, no filtering by finished status is applied.
func (l *Library) GetBooks(author string, isFinished *bool) map[string]Book {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Book, len(l.books))

	for title, book := range l.books {
		if author != "" && author != book.Author {
			continue
		}
		if isFinished != nil && *isFinished != book.IsFinished {
			continue
		}

		tmp[title] = book
	}

	return tmp
}

// DeleteBook deletes the book with given title from library.
//
// It returns ErrBookNotFound if the book does not exist.
func (l *Library) DeleteBook(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.books[title]; !ok {
		return ErrBookNotFound
	}

	delete(l.books, title)

	return nil
}
