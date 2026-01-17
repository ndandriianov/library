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
	str, err := json.MarshalIndent(e, "", "\n")
	if err != nil {
		// If marshaling fails, return a valid JSON error message
		// This prevents error handling from failing silently
		return fmt.Sprintf(`{"Message": "failed to marshal error: %v", "Time": "%s"}`, err, e.Time.Format(time.RFC3339))
	}

	return string(str)
}
