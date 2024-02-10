package errortypes

import "fmt"

type ApiError struct {
	Code    int
	Message string
	Data    any
}

func (err *ApiError) Error() string {
	return fmt.Sprintf("ERROR %d: %s", err.Code, err.Message)
}
