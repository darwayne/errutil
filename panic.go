package errutil

import (
	"runtime"

	"errors"
)

// OnExpectedPanic recovers from an expected panic and provides expected error to the provided callback
// if panic is NOT an error interface it will be ra-raised
// example:
// defer OnExpected(func(err error){ doSomethingWithError(err) })
//
// note: this should be used sparingly; public APIs should NOT expect consumers to handle panics
func OnExpectedPanic(fn func(err error)) {
	r := recover()
	var runtimeErr runtime.Error
	if err, ok := r.(error); ok && !errors.As(err, &runtimeErr) {
		fn(err)
	} else if r != nil {
		panic(r)
	}
}

// ExpectedPanicAsError sets the error pointer with the expected panic
// if panic is NOT an error interface it will be ra-raised
// note: this function must be deferred e.g.
// defer ExpectedPanicAsErr(&myErr)
//
// note: this should be used sparingly; public APIs should NOT expect consumers to handle panics
func ExpectedPanicAsError(errPtr *error) {
	r := recover()
	var runtimeErr runtime.Error
	if err, ok := r.(error); ok && !errors.As(err, &runtimeErr) {
		*errPtr = err
	} else if r != nil {
		panic(r)
	}
}
