package api

import "fmt"


type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("status code: %d with message: %s\n", e.StatusCode, e.Message)
}
