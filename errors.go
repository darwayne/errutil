package errutil

// New creates a new error with the ability to add multiple behavior to that error
// e.g.
//
// New(errors.New("my error"), WithAccessDenied(true), WithRateLimit(true))
func New(err error, opts ...OptsFunc) error {
	e := multiKindErr{
		opts:  ToOpts(opts...),
		error: err,
	}

	return e
}

var (
	_ AccessDenier = (*multiKindErr)(nil)
	_ Conflicter   = (*multiKindErr)(nil)
	_ Exister      = (*multiKindErr)(nil)
	_ NotFounder   = (*multiKindErr)(nil)
	_ RateLimiter  = (*multiKindErr)(nil)
	_ TooLarge     = (*multiKindErr)(nil)
	_ error        = (*multiKindErr)(nil)
)

type Opts struct {
	//::builder-gen -no-builder -with-globals -prefix=With -with-optional-bools
	AccessDenied bool
	Conflict     bool
	NotFound     bool
	Exists       bool
	TooLarge     bool
	RateLimit    bool
}

type multiKindErr struct {
	opts Opts

	error
}

func (m multiKindErr) AccessDenied() bool {
	return m.opts.AccessDenied
}

func (m multiKindErr) Conflict() bool {
	return m.opts.Conflict
}

func (m multiKindErr) Exists() bool {
	return m.opts.Exists
}

func (m multiKindErr) NotFound() bool {
	return m.opts.NotFound
}

func (m multiKindErr) RateLimit() bool {
	return m.opts.RateLimit
}

func (m multiKindErr) TooLarge() bool {
	return m.opts.TooLarge
}

func (m multiKindErr) Unwrap() error {
	return m.error
}
