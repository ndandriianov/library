package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{handlers: handlers}
}

func (s Server) Serve() error {
	router := mux.NewRouter()

	router.Use(JsonMiddleware)

	router.Path("/books").Methods("POST").HandlerFunc(s.handlers.HandleAddBook)
	router.Path("/books/{title}/finish").Methods("PATCH").HandlerFunc(s.handlers.HandleFinishBook)
	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.handlers.HandleGetBook)
	router.Path("/books").Methods("GET").HandlerFunc(s.handlers.HandleGetBooks)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.handlers.HandleDeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
