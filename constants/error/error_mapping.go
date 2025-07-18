package error

import "errors"

func ErrMapping(err error) bool {
	allErrors := append(GeneralError, UserErrors...)

	for _, item := range allErrors {
		if errors.Is(err, item) {
			return true
		}
	}
	return false
}
