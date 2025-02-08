package apperror

import "net/http"

type UnauthorizedError struct {
	errmsg string
}

var _ AppError = UnauthorizedError{}

func (i UnauthorizedError) Code() int {
	return http.StatusUnauthorized
}

func (i UnauthorizedError) Message() string {
	return i.errmsg
}

func (i UnauthorizedError) Cause() any {
	return nil
}

func (i UnauthorizedError) ShouldLog() bool {
	return false
}

func ErrUnauthorized(errmsg string) UnauthorizedError {
	return UnauthorizedError{
		errmsg,
	}
}
