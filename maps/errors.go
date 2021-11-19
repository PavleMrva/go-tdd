package maps

import "errors"

var (
	ErrNotFound = errors.New(notFoundErrorMessage)
	ErrWordExists = errors.New(existsErrorMessage)
	ErrWordDoesNotExist = errors.New(doesNotExistErrorMessage)
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}