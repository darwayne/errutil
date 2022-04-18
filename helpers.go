package errutil

// NewAccessDenied creates a new error and gives it AccessDenier behavior
func NewAccessDenied(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithAccessDenied())
	return New(msg, opts...)
}

// WrapAccessDenied wraps an error and gives it AccessDenier behavior
func WrapAccessDenied(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithAccessDenied())
	return Wrap(err, opts...)
}

// NewExists creates a new error and gives it Exister behavior
func NewExists(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithExists())
	return New(msg, opts...)
}

// WrapExists wraps an error and gives it Exister behavior
func WrapExists(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithExists())
	return Wrap(err, opts...)
}

// NewConflict creates a new error and gives it Conflicter behavior
func NewConflict(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithConflict())
	return New(msg, opts...)
}

// WrapConflict wraps an error and gives it Conflicter behavior
func WrapConflict(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithConflict())
	return Wrap(err, opts...)
}

// NewNotFound creates a new error and gives it NotFounder behavior
func NewNotFound(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithNotFound())
	return New(msg, opts...)
}

// WrapNotFound wraps an error and gives it NotFounder behavior
func WrapNotFound(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithNotFound())
	return Wrap(err, opts...)
}

// NewRateLimit creates a new error and gives it RateLimiter behavior
func NewRateLimit(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithRateLimit())
	return New(msg, opts...)
}

// WrapRateLimit wraps an error and gives it RateLimiter behavior
func WrapRateLimit(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithRateLimit())
	return Wrap(err, opts...)
}

// NewTemporary creates a new error and gives it Temporarily behavior
func NewTemporary(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTemporary())
	return New(msg, opts...)
}

// WrapTemporary wraps an error and gives it Temporarily behavior
func WrapTemporary(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTemporary())
	return Wrap(err, opts...)
}

// NewTooLarge creates a new error and gives it TooLarge behavior
func NewTooLarge(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTooLarge())
	return New(msg, opts...)
}

// WrapTooLarge wraps an error and gives it TooLarge behavior
func WrapTooLarge(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTooLarge())
	return Wrap(err, opts...)
}

// NewTooManyRequests creates a new error and gives it TooManyRequester behavior
func NewTooManyRequests(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTooManyRequests())
	return New(msg, opts...)
}

// WrapTooManyRequests wraps an error and gives it TooManyRequester behavior
func WrapTooManyRequests(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithTooManyRequests())
	return Wrap(err, opts...)
}

// NewUnauthorized creates a new error and gives it Unauthorizable behavior
func NewUnauthorized(msg string, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithUnauthorized())
	return New(msg, opts...)
}

// WrapUnauthorized wraps an error and gives it Unauthorizable behavior
func WrapUnauthorized(err error, opts ...OptsFunc) error {
	opts = append(opts, WithStackTrace(1), WithUnauthorized())
	return Wrap(err, opts...)
}
