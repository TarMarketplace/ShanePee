package apperror

import "net/http"

type InternalError struct {
	cause any
}

var _ AppError = InternalError{}

func (i InternalError) Code() int {
	return http.StatusInternalServerError
}

func (i InternalError) Message() string {
	return "Internal server error"
}

func (i InternalError) Cause() any {
	return i.cause
}

func (i InternalError) ShouldLog() bool {
	return true
}

func ErrInternal(cause any) InternalError {
	return InternalError{
		cause,
	}
}
