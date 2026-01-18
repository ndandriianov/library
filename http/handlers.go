package http

import (
	"encoding/json"
	"errors"
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

	if err := json.NewEncoder(w).Encode(book); err != nil {
		logFailedWriteHTTPResponse(err)
		return
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

	if err := json.NewEncoder(w).Encode(book); err != nil {
		logFailedWriteHTTPResponse(err)
		return
	}
}

func (h *Handlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	book, err := h.lib.GetBook(title)
	if err != nil {
		writeError(w, err, http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		logFailedWriteHTTPResponse(err)
		return
	}
}

func (h *Handlers) HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	author := q.Get("author")
	isFinishedStr := q.Get("isFinished")

	var isFinished *bool
	if isFinishedStr == "true" {
		tmp := true
		isFinished = &tmp
	} else if isFinishedStr == "false" {
		tmp := false
		isFinished = &tmp
	}

	books := h.lib.GetBooks(author, isFinished)
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		logFailedWriteHTTPResponse(err)
		return
	}
}

func (h *Handlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.lib.DeleteBook(title); err != nil {
		switch {
		case errors.Is(err, library.ErrBookNotFound):
			writeError(w, err, http.StatusNotFound)
		default:
			writeError(w, err, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
