package errutil

// StringErr allows you to use a string as an error
type StringErr string

func (s StringErr) Error() string {
	return string(s)
}
