package main

import (
	"LibraryManager/http"
	"LibraryManager/library"
	"fmt"
)

func main() {
	lib := library.NewLibrary()
	handlers := http.NewHandlers(&lib)
	server := http.NewServer(handlers)

	if err := server.Serve(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
