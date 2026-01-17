package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"library/http/dto"
	"library/library"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	lib *library.Library
}

func (h *Handlers) HandleAddBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO dto.Book
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	book, err := library.NewBook(bookDTO.Title, bookDTO.Author, bookDTO.NumberOfPages)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.lib.AddBook(book); err != nil {
		writeError(w, err, http.StatusConflict)
		return
	}

	b, err := json.Marshal(book)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
	}
}

func (h *Handlers) HandleFinishBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	book, err := h.lib.FinishBook(title)
	if err != nil {
		var status int
		switch {
		case errors.Is(err, library.ErrBookNotFound):
			status = http.StatusNotFound
		case errors.Is(err, library.ErrBookIsAlreadyFinished):
			status = http.StatusConflict
		default:
			status = http.StatusInternalServerError
		}
		writeError(w, err, status)
		return
	}

	b, err := json.Marshal(book)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
	}
}
