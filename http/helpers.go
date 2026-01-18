package http

import (
	"fmt"
	"library/http/dto"
	"net/http"
	"time"
)

func writeError(w http.ResponseWriter, currentError error, status int) {
	w.WriteHeader(status)

	errDTO := dto.Err{
		Message: currentError.Error(),
		Time:    time.Now(),
	}

	_, err := w.Write([]byte(errDTO.ToString()))
	if err != nil {
		fmt.Println("failed to write http response")
	}
}

func logFailedWriteHTTPResponse(err error) {
	fmt.Println("failed to write HTTP response", err)
}
