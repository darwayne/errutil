package errutil

// NewCode adds code to an error
func NewCode(err error, code string) error {
	return coded{
		wrappedError: wrappedError{error: err},
		code:         code,
	}
}

type coded struct {
	wrappedError
	code string
}

func (c coded) Code() string {
	return c.code
}
