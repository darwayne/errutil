package errutil

import (
	"testing"
)

func TestExpectedPanicAsError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		result := func() (e error) {
			defer ExpectedPanicAsError(&e)
			panic(StringErr("oh snap"))
			return
		}()

		if result == nil || result != StringErr("oh snap") {
			t.Fatal("ruh oh")
		}
	})

	t.Run("should re-raise panic", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected to recover")
			}
			val, ok := r.(int)
			if !ok {
				t.Fatalf("expected an int but got %T", r)
			}
			if val != 42 {
				t.Fatal("expected to get 42 but got", val)
			}
		}()

		_ = func() (e error) {
			defer ExpectedPanicAsError(&e)
			panic(42)
			return
		}()
	})
}

func TestOnExpectedPanic(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		var called bool
		result := func() (e error) {
			defer OnExpectedPanic(func(err error) {
				called = true
				e = err
			})
			panic(StringErr("oh snap"))
			return
		}()

		if result == nil || result != StringErr("oh snap") {
			t.Fatal("ruh oh")
		}
		if !called {
			t.Fatal("function not called as expected")
		}
	})

	t.Run("should re-raise panic", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected to recover")
			}
			val, ok := r.(int)
			if !ok {
				t.Fatalf("expected an int but got %T", r)
			}
			if val != 42 {
				t.Fatal("expected to get 42 but got", val)
			}
		}()

		_ = func() (e error) {
			defer OnExpectedPanic(func(err error) {
				e = err
			})
			panic(42)
			return
		}()
	})
}
