package http

import (
	"encoding/json"
	"fmt"
	"library/http/dto"
	"library/library"
	"net/http"
	"time"
)

type Handlers struct {
	lib *library.Library
}

func (h *Handlers) HandleAddBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO dto.Book
	if err := json.NewDecoder(r.Body).Decode(bookDTO); err != nil {
		errDTO := dto.Err{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	book, err := library.NewBook(bookDTO.Title, bookDTO.Author, bookDTO.NumberOfPages)
	if err != nil {
		errDTO := dto.Err{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	h.lib.AddBook(book)

	b, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		errDTO := dto.Err{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
	}
}

//func (h *Handlers) WriteError(w http.ResponseWriter, err error, rules map[error]int) {
//
//}
