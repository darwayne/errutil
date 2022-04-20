package errutil

import (
	"fmt"
	"io"
	"runtime"

	"github.com/pkg/errors"
)

func NewStacked(err error, skip ...int) error {
	if err == nil {
		return nil
	}

	return &Stacked{
		error: err,
		stack: Callers(skip...),
	}
}

type Stacked struct {
	error
	*stack
}

func (w *Stacked) Unwrap() error { return w.error }

func (w *Stacked) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.error)
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := errors.Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*s)[i])
	}
	return f
}

func Callers(skip ...int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	s := 0
	if len(skip) > 0 {
		s = skip[0]
	}

	n := runtime.Callers(s+3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
