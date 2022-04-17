package errutil

import (
	"errors"
)

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

// IsTemporary checks if an error exhibits Temporary behavior
func IsTemporary(err error) bool {
	if err == nil {
		return false
	}

	var e Temporarily
	if errors.As(err, &e) && e.Temporary() {
		return true
	}

	return IsTemporary(errors.Unwrap(err))
}

// IsStatusCodeable checks if an error exhibits StatusCodeable behavior
func IsStatusCodeable(err error) bool {
	var e StatusCodeable
	if err == nil || !errors.As(err, &e) {
		return false
	}

	return true
}

// IsCodeable checks if an error exhibits Codeable behavior
func IsCodeable(err error) bool {
	var e Codeable
	if err == nil || !errors.As(err, &e) {
		return false
	}

	return true
}

// GetStatusCode returns the status code for an error if it has one
func GetStatusCode(err error) int {
	var e StatusCodeable
	if err == nil || !errors.As(err, &e) {
		return 0
	}

	return e.StatusCode()
}

// GetCode returns the code for an error if it has one
func GetCode(err error) string {
	var e Codeable
	if err == nil || !errors.As(err, &e) {
		return ""
	}

	return e.Code()
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

// IsTooManyRequests checks if an error exhibits TooManyRequester behavior
func IsTooManyRequests(err error) bool {
	if err == nil {
		return false
	}

	var e TooManyRequester
	if errors.As(err, &e) && e.TooManyRequests() {
		return true
	}

	return IsTooManyRequests(errors.Unwrap(err))
}

// IsInternalErrorMessage checks if an error exhibits InternalErrorMessagable behavior
func IsInternalErrorMessage(err error) bool {
	var e InternalErrorMessagable
	if err == nil || errors.As(err, &e) {
		return false
	}

	return true
}

// GetInternalErrorMessage returns the internal error message for an error if it has one
func GetInternalErrorMessage(err error) string {
	var e InternalErrorMessagable
	if err == nil || !errors.As(err, &e) {
		return ""
	}

	return e.InternalErrorMessage()
}

// IsTaggable checks if an error exhibits taggable behavior
func IsTaggable(err error) bool {
	if err == nil {
		return false
	}

	var e Taggable
	if errors.As(err, &e) && len(e.Tags()) > 0 {
		return true
	}

	return IsTaggable(errors.Unwrap(err))
}

// IsStackTraceable checks if an error exhibits taggable behavior
func IsStackTraceable(err error) bool {
	if err == nil {
		return false
	}

	var e StackTraceable
	if errors.As(err, &e) {
		return true
	}

	return IsStackTraceable(errors.Unwrap(err))
}

// IsUnauthorized checks if an error exhibits Unauthorized behavior
func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}

	var e Unauthorizable
	if errors.As(err, &e) && e.Unauthorized() {
		return true
	}

	return IsUnauthorized(errors.Unwrap(err))
}

// GetTags returns all the tags for a given error
func GetTags(err error) []Tag {
	var tags []Tag
	for {
		if err == nil {
			break
		}
		var e Tagged
		if errors.As(err, &e) && len(e.Tags()) > 0 {
			tags = append(tags, e.Tags()...)
		}

		err = errors.Unwrap(e)
	}

	return tags
}

func WithEasyTags(key, value string, additionalKvs ...string) OptsFunc {
	return func(opts *Opts) {
		tags := make([]Tag, 0, 10)
		tags = append(tags, Tag{Key: key, Value: value})
		if len(additionalKvs) > 0 && len(additionalKvs)%2 == 0 {
			for i := 0; i < len(additionalKvs); i = i + 2 {
				tags = append(tags, Tag{Key: additionalKvs[i], Value: additionalKvs[i+1]})
			}
		}

		opts.Tags = tags
	}
}
