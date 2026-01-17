package http

import (
	"encoding/json"
	"fmt"
	"library/http/dto"
	"library/library"
	"net/http"
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

	b, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
	}
}
