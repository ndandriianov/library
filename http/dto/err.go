package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type Err struct {
	Message string
	Time    time.Time
}

func (e *Err) ToString() string {
	str, err := json.Marshal(e)
	if err != nil {
		msg := "failed to marshall Err"
		fmt.Println(msg)
		return msg
	}

	return string(str)
}
