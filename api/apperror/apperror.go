package apperror

type AppError interface {
	// Public facing messing
	Message() string

	// HTTP Status code
	Code() int

	// Internal logging message, return nil if it should not be log
	Cause() any

	ShouldLog() bool
}
