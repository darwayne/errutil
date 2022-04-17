package errutil

import "github.com/pkg/errors"

// AccessDenier determines if an error exhibits access denied behavior
type AccessDenier interface {
	AccessDenied() bool
}

// Exister determines if an error exhibits exists behavior
type Exister interface {
	Exists() bool
}

// NotFounder determines if an error exhibits not found behavior
type NotFounder interface {
	NotFound() bool
}

// RateLimiter determines if an error exhibits rate limit behavior
type RateLimiter interface {
	RateLimit() bool
}

// Conflicter determines if an error exhibits conflict behavior
type Conflicter interface {
	Conflict() bool
}

// TooLarge determines if an error exhibits too large behavior
type TooLarge interface {
	TooLarge() bool
}

// Taggable determines if an error exhibits tag behavior
type Taggable interface {
	Tags() []Tag
}

// StackTraceable determines if an error exhibits stacktrace behavior from the pkg/errors package
type StackTraceable interface {
	StackTrace() errors.StackTrace
}

// Temporarily determines if an error exhibits temporary behavior
type Temporarily interface {
	Temporary() bool
}

type TooManyRequester interface {
	TooManyRequests() bool
}

type StatusCodeable interface {
	StatusCode() int
}

type Codeable interface {
	Code() string
}

type InternalErrorMessagable interface {
	InternalErrorMessage() string
}

// Unauthorizable determines if an error exhibits temporary behavior
type Unauthorizable interface {
	Unauthorized() bool
}
