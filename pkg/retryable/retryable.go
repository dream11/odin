package retryable

type Error struct {
	err       error
	retryable bool
}

func (w Error) Error() string {
	return w.err.Error()
}

func (w Error) Retryable() bool {
	return w.retryable
}

func NewRetryableError(err error, retryable bool) Error {
	return Error{
		err:       err,
		retryable: retryable,
	}
}
