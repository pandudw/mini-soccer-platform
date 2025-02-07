package error

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordIncorrect  = errors.New("password incorrect")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrPasswordNotMatch   = errors.New("password not match")
)
var UserError = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUserAlreadyExists,
	ErrEmailAlreadyExists,
	ErrPasswordNotMatch,
}
