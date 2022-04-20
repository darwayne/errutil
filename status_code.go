package errutil

// NewStatusCode adds a status code to an error
func NewStatusCode(err error, code int) error {
	if err == nil {
		return nil
	}

	return statusCode{
		wrappedError: wrappedError{error: err},
		code:         code,
	}
}

type statusCode struct {
	wrappedError
	code int
}

func (i statusCode) StatusCode() int {
	return i.code
}
