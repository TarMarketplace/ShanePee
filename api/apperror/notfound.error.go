package apperror

import "net/http"

type NotFoundError struct {
	errmsg string
}

var _ AppError = NotFoundError{}

func (i NotFoundError) Code() int {
	return http.StatusNotFound
}

func (i NotFoundError) Message() string {
	return i.errmsg
}

func (i NotFoundError) Cause() any {
	return nil
}

func (i NotFoundError) ShouldLog() bool {
	return false
}

func ErrNotFound(errmsg string) NotFoundError {
	return NotFoundError{
		errmsg,
	}
}
