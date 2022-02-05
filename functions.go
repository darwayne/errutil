package errutil

import "errors"

type CheckerFn func(error) bool

// IsAccessDenied checks if an error exhibits AccessDenier behavior
func IsAccessDenied(err error) bool {
	if err == nil {
		return false
	}

	var e AccessDenier
	if errors.As(err, &e) && e.AccessDenied() {
		return true
	}

	return IsAccessDenied(errors.Unwrap(err))
}

// IsConflict checks if an error exhibits Conflicter behavior
func IsConflict(err error) bool {
	if err == nil {
		return false
	}

	var e Conflicter
	if errors.As(err, &e) && e.Conflict() {
		return true
	}

	return IsConflict(errors.Unwrap(err))
}

// IsExist checks if an error exhibits Exister behavior
func IsExist(err error) bool {
	if err == nil {
		return false
	}

	var e Exister
	if errors.As(err, &e) && e.Exists() {
		return true
	}

	return IsExist(errors.Unwrap(err))
}

// IsNotFound checks if an error exhibits NotFounder behavior
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	var e NotFounder
	if errors.As(err, &e) && e.NotFound() {
		return true
	}

	return IsNotFound(errors.Unwrap(err))
}

// IsRateLimit checks if an error exhibits RateLimiter behavior
func IsRateLimit(err error) bool {
	if err == nil {
		return false
	}

	var e RateLimiter
	if errors.As(err, &e) && e.RateLimit() {
		return true
	}

	return IsRateLimit(errors.Unwrap(err))
}

// IsTooLarge checks if an error exhibits TooLarger behavior
func IsTooLarge(err error) bool {
	if err == nil {
		return false
	}

	var e TooLarge
	if errors.As(err, &e) && e.TooLarge() {
		return true
	}

	return IsTooLarge(errors.Unwrap(err))
}
