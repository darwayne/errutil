package errutil

import (
	"errors"
	"fmt"
	"strings"
)

// Wrap wraps an error with the ability to add multiple behavior to that error
// e.g.
//
// Wrap(errors.New("my error"), WithAccessDenied(true), WithRateLimit(true))
func Wrap(err error, opts ...OptsFunc) error {
	if err == nil {
		return nil
	}

	e := multiKindErr{
		opts:  ToOpts(opts...),
		error: err,
	}

	if len(e.opts.Tags) > 0 {
		e.error = NewTagged(err, e.opts.Tags...)
	}
	if e.opts.StatusCode != nil {
		e.error = NewStatusCode(e.error, *e.opts.StatusCode)
	}
	if e.opts.InternalErrorMessage != nil {
		e.error = NewInternalErrorMessage(e.error, *e.opts.InternalErrorMessage)
	}
	if e.opts.Code != nil {
		e.error = NewCode(e.error, *e.opts.Code)
	}
	if e.opts.StackTrace != nil {
		e.error = NewStacked(e.error, *e.opts.StackTrace+1)
	} else {
		e.error = NewStacked(e.error, 1)
	}

	return e
}

// New creates a new error with the ability to add multiple behavior to that error
// e.g.
//
// New("my error", WithAccessDenied(true), WithRateLimit(true))
func New(message string, opts ...OptsFunc) error {
	info := ToOpts(opts...)
	if info.StackTrace != nil {
		opts = append(opts, WithStackTrace(*info.StackTrace+1))
	} else {
		opts = append(opts, WithStackTrace(1))
	}

	return Wrap(errors.New(message), opts...)
}

var (
	_ AccessDenier     = (*multiKindErr)(nil)
	_ Conflicter       = (*multiKindErr)(nil)
	_ Exister          = (*multiKindErr)(nil)
	_ NotFounder       = (*multiKindErr)(nil)
	_ RateLimiter      = (*multiKindErr)(nil)
	_ TooLarge         = (*multiKindErr)(nil)
	_ TooManyRequester = (*multiKindErr)(nil)
	_ Temporarily      = (*multiKindErr)(nil)
	_ Unauthorizable   = (*multiKindErr)(nil)
	_ error            = (*multiKindErr)(nil)
)

type Opts struct {
	//::builder-gen -no-builder -with-globals -prefix=With -with-optional-bools
	AccessDenied         bool
	Conflict             bool
	NotFound             bool
	Exists               bool
	TooLarge             bool
	RateLimit            bool
	Tags                 []Tag
	StackTrace           *int
	Unauthorized         bool
	Temporary            bool
	InternalErrorMessage *string
	Code                 *string
	StatusCode           *int
	TooManyRequests      bool
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

func (m multiKindErr) Format(s fmt.State, verb rune) {
	if formatter, ok := m.error.(fmt.Formatter); ok {
		formatter.Format(s, verb)
		return
	} else {
		var builder strings.Builder
		builder.WriteString("%")
		if s.Flag('+') {
			builder.WriteRune('+')
		}
		builder.WriteRune(verb)
		fmt.Fprintf(s, builder.String(), m.error)
	}
}

func (m multiKindErr) NotFound() bool {
	return m.opts.NotFound
}

func (m multiKindErr) RateLimit() bool {
	return m.opts.RateLimit
}

func (m multiKindErr) Temporary() bool {
	return m.opts.Temporary
}

func (m multiKindErr) TooLarge() bool {
	return m.opts.TooLarge
}

func (m multiKindErr) TooManyRequests() bool {
	return m.opts.TooManyRequests
}

func (m multiKindErr) Unauthorized() bool {
	return m.opts.Unauthorized
}

func (m multiKindErr) Unwrap() error {
	return m.error
}
