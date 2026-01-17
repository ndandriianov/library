package library

import "errors"

var ErrBookAlreadyExists = errors.New("book already exists")
var ErrBookNotFound = errors.New("book not found")
var ErrBookIsAlreadyFinished = errors.New("book is already finished")
var ErrInvalidBookTitle = errors.New("book title is invalid")
var ErrInvalidBookAuthor = errors.New("book author is invalid")
var ErrInvalidBookNumberOfPages = errors.New("book number of pages is invalid")
