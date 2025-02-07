package error

import "errors"

var (
	ErrInternalServer  = errors.New("internal server error")
	ErrDatabase        = errors.New("database error")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidToken    = errors.New("invalid token")
	ErrForbidden       = errors.New("forbidden")
)

var GeneralErrors = []error{
	ErrInternalServer,
	ErrDatabase,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
