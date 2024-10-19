package cmdUtils

import "fmt"

// A *flagError indicates an error processing command-line flags or other arguments.
// Such errors cause the application to display the usage message.
type flagError struct {
	// Note: not struct{error}: only *FlagError should satisfy error.
	err error
}

func (fe *flagError) Error() string {
	return fe.err.Error()
}

// FlagErrorf returns a new FlagError that wraps an error produced by
// fmt.Errorf(format, args...).
func FlagErrorf(format string, args ...interface{}) error {
	return flagErrorWrap(fmt.Errorf(format, args...))
}

// FlagError returns a new FlagError that wraps the specified error.
func flagErrorWrap(err error) error { return &flagError{err} }
