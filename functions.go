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
