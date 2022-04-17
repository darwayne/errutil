package errutil

func NewInternalErrorMessage(err error, msg string) error {
	return internalErrorMessage{
		wrappedError: wrappedError{error: err},
		msg:          msg,
	}
}

type internalErrorMessage struct {
	wrappedError
	msg string
}

func (i internalErrorMessage) InternalErrorMessage() string {
	return i.msg
}
