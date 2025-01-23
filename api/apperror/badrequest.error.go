package apperror

import "net/http"

type BadRequestError struct {
	errmsg string
}

var _ AppError = BadRequestError{}

func (i BadRequestError) Code() int {
	return http.StatusInternalServerError
}

func (i BadRequestError) Message() string {
	return "Internal server error"
}

func (i BadRequestError) Cause() any {
	return nil
}

func (i BadRequestError) ShouldLog() bool {
	return false
}

func ErrBadRequest(errmsg string) InternalError {
	return InternalError{
		errmsg,
	}
}
