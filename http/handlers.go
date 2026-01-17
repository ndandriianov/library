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
	if err := json.NewDecoder(r.Body).Decode(&bookDTO); err != nil {
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

	if err := h.lib.AddBook(book); err != nil {
		errDTO := dto.Err{
			Message: err.Error(),
			Time:    time.Now(),
		}

		statusCode := http.StatusInternalServerError
		if err == library.ErrBookAlreadyExists {
			statusCode = http.StatusConflict
		}

		http.Error(w, errDTO.ToString(), statusCode)
		return
	}

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
		// Log error but don't return error response as headers are already sent
		errDTO := dto.Err{
			Message: fmt.Sprintf("failed to write http response: %v", err),
			Time:    time.Now(),
		}
		fmt.Fprintf(w, "%s\n", errDTO.ToString())
	}
}
