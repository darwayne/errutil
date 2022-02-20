// Code generated by builder-gen. DO NOT EDIT.
package errutil

type OptsFunc func(*Opts)

func ToOpts(opts ...OptsFunc) Opts {
	var info Opts
	ToOptsWithDefault(&info, opts...)

	return info
}

func ToOptsWithDefault(info *Opts, opts ...OptsFunc) {
	for _, o := range opts {
		o(info)
	}
}

func WithAccessDenied(accessDeniedParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.AccessDenied = len(accessDeniedParam) == 0 || accessDeniedParam[0]
	}
}

func WithConflict(conflictParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.Conflict = len(conflictParam) == 0 || conflictParam[0]
	}
}

func WithNotFound(notFoundParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.NotFound = len(notFoundParam) == 0 || notFoundParam[0]
	}
}

func WithExists(existsParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.Exists = len(existsParam) == 0 || existsParam[0]
	}
}

func WithTooLarge(tooLargeParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.TooLarge = len(tooLargeParam) == 0 || tooLargeParam[0]
	}
}

func WithRateLimit(rateLimitParam ...bool) OptsFunc {
	return func(opts *Opts) {
		opts.RateLimit = len(rateLimitParam) == 0 || rateLimitParam[0]
	}
}