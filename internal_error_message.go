package errutil

// NewInternalErrorMessage adds an internal message to an error
func NewInternalErrorMessage(err error, msg string) error {
	if err == nil {
		return nil
	}

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
