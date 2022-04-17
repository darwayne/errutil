package errutil

func NewStatusCode(err error, code int) error {
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
