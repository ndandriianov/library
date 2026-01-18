package http

import (
	"LibraryManager/http/dto"
	"fmt"
	"net/http"
	"time"
)

func writeError(w http.ResponseWriter, currentError error, status int) {
	w.WriteHeader(status)

	errDTO := dto.Err{
		Message: currentError.Error(),
		Time:    time.Now(),
	}

	if _, err := w.Write([]byte(errDTO.ToString())); err != nil {
		logFailedWriteHTTPResponse(err)
		return
	}
}

func logFailedWriteHTTPResponse(err error) {
	fmt.Println("failed to write HTTP response", err)
}
