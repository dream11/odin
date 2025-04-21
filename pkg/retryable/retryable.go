package retryable

// Error Package retryable provides a way to handle errors that can be retried.
type Error struct {
	err       error
	retryable bool
}

// Error creates a new retryable error.
func (w Error) Error() string {
	return w.err.Error()
}

// Retryable checks if the error is retryable.
func (w Error) Retryable() bool {
	return w.retryable
}

// NewRetryableError creates a new retryable error.
func NewRetryableError(err error, retryable bool) Error {
	return Error{
		err:       err,
		retryable: retryable,
	}
}
