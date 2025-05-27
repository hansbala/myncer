package core

import "fmt"

func WrappedError(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func NewError(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
