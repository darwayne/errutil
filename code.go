package errutil

func NewCode(err error, msg string) error {
	return code{
		wrappedError: wrappedError{error: err},
		msg:          msg,
	}
}

type code struct {
	wrappedError
	msg string
}

func (c code) Code() string {
	return c.msg
}
