package errutil

type wrappedError struct {
	error
}

func (u wrappedError) Unwrap() error {
	return u.error
}
