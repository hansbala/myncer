package core

import "fmt"

func WrappedError(err error, msg string) error {
	return fmt.Errorf("msg: %w", err)
}
