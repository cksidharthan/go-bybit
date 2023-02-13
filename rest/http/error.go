package http

import "fmt"

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("status code: %d | message: %s", e.StatusCode, e.Message)
}
