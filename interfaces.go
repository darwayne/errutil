package errutil

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
