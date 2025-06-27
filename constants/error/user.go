package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordDoesNotMatch = errors.New("password doesn't match")
	ErrPasswordIncorrect    = errors.New("incorrect password")
	ErrUsernameExist        = errors.New("username exist")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordDoesNotMatch,
	ErrPasswordIncorrect,
	ErrUsernameExist,
}
