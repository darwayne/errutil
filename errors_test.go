package errutil

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestWrapError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		err := errors.New("yo son")

		t.Run("shared", func(t *testing.T) {
			for idx, tt := range sharedHappyPathTests {
				t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
					e := Wrap(err, tt.opts...)
					if len(tt.fns) != len(tt.expectations) {
						t.Fatal("fns and expectations should be the same length")
					}

					for idx, fn := range tt.fns {
						if tt.expectations[idx] != fn(e) {
							t.Fatalf("expected %s to return %v", getFunctionName(fn), tt.expectations[idx])
						}
					}
				})
			}
		})

		t.Run("with internal error message", func(t *testing.T) {
			e := Wrap(err, WithInternalErrorMessage("yolo"))
			if !IsInternalErrorMessage(e) {
				t.Fatal("expected true")
			}

			if GetInternalErrorMessage(e) != "yolo" {
				t.Fatal("expected yolo")
			}
		})

		t.Run("with code", func(t *testing.T) {
			e := Wrap(err, WithCode("yolo"))
			if !IsCodeable(e) {
				t.Fatal("expected true")
			}

			if GetCode(e) != "yolo" {
				t.Fatal("expected yolo")
			}
		})

		t.Run("with status code", func(t *testing.T) {
			e := Wrap(err, WithStatusCode(23))
			if !IsStatusCodeable(e) {
				t.Fatal("expected true")
			}

			if GetStatusCode(e) != 23 {
				t.Fatal("expected 23")
			}
		})

		t.Run("with too many requests", func(t *testing.T) {
			e := Wrap(err, WithTooManyRequests())
			if !IsTooManyRequests(e) {
				t.Fatal("expected true")
			}
		})

		t.Run("with nil error", func(t *testing.T) {
			err := Wrap(nil, WithNotFound())
			if err != nil {
				t.Fatal("expected nil")
			}
		})

	})

	t.Run("should handle tagged errors as expected", func(t *testing.T) {
		e := Wrap(errors.New("hi"), WithTags(NewTag("what", "is"), NewTag("hello", "world")))
		if !IsTaggable(e) {
			t.Fatal("expected IsTaggable to return true")
		}

		if tags := GetTags(e); len(tags) != 2 {
			t.Fatalf("expected 2 tags back but got %d", len(tags))
		} else {
			t.Log(tags)
		}

		t.Run("should handle nested tags as well", func(t *testing.T) {
			nested := Wrap(e, WithEasyTags("yes", "sir", "will", "be"))
			if !IsTaggable(nested) {
				t.Fatal("expected IsTaggable to return true")
			}

			if tags := GetTags(nested); len(tags) != 4 {
				t.Fatalf("expected 4 tags back but got %d", len(tags))
			} else {
				t.Log(tags)
			}
		})

	})

	t.Run("should generate stack trace when option provided", func(t *testing.T) {
		err := Wrap(errors.New("hey there"), WithStackTrace(0))
		fmt.Printf("%+v\n", err)
	})
}
func TestNewError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		err := "yo son"

		for idx, tt := range sharedHappyPathTests {
			t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
				e := New(err, tt.opts...)
				if len(tt.fns) != len(tt.expectations) {
					t.Fatal("fns and expectations should be the same length")
				}

				for idx, fn := range tt.fns {
					if tt.expectations[idx] != fn(e) {
						t.Fatalf("expected %s to return %v", getFunctionName(fn), tt.expectations[idx])
					}
				}
			})
		}
	})

	t.Run("should handle tagged errors as expected", func(t *testing.T) {
		e := New("hi", WithTags(NewTag("what", "is"), NewTag("hello", "world")))
		if !IsTaggable(e) {
			t.Fatal("expected IsTaggable to return true")
		}

		if tags := GetTags(e); len(tags) != 2 {
			t.Fatalf("expected 2 tags back but got %d", len(tags))
		} else {
			t.Log(tags)
		}

		t.Run("should handle nested tags as well", func(t *testing.T) {
			nested := Wrap(e, WithEasyTags("yes", "sir", "will", "be"))
			if !IsTaggable(nested) {
				t.Fatal("expected IsTaggable to return true")
			}

			if tags := GetTags(nested); len(tags) != 4 {
				t.Fatalf("expected 4 tags back but got %d", len(tags))
			} else {
				t.Log(tags)
			}
		})

	})

	t.Run("should generate stack trace when option provided", func(t *testing.T) {
		err := New("hey there", WithStackTrace(0))
		fmt.Printf("%+v\n", err)
	})
}

var sharedHappyPathTests = []struct {
	opts         []OptsFunc
	fns          []CheckerFn
	expectations []bool
}{
	{
		opts:         []OptsFunc{WithAccessDenied(true)},
		fns:          append([]CheckerFn{IsAccessDenied}, allFuncsExcept(IsAccessDenied)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts:         []OptsFunc{WithConflict(true)},
		fns:          append([]CheckerFn{IsConflict}, allFuncsExcept(IsConflict)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts:         []OptsFunc{WithExists(true)},
		fns:          append([]CheckerFn{IsExist}, allFuncsExcept(IsExist)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts:         []OptsFunc{WithNotFound(true)},
		fns:          append([]CheckerFn{IsNotFound}, allFuncsExcept(IsNotFound)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts:         []OptsFunc{WithRateLimit(true)},
		fns:          append([]CheckerFn{IsRateLimit}, allFuncsExcept(IsRateLimit)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts:         []OptsFunc{WithTooLarge(true)},
		fns:          append([]CheckerFn{IsTooLarge}, allFuncsExcept(IsTooLarge)...),
		expectations: []bool{true, false, false, false, false, false, false, false},
	},
	{
		opts: []OptsFunc{WithAccessDenied(true), WithTooLarge(true),
			WithConflict(true), WithNotFound(true), WithExists(true),
			WithRateLimit(true), WithTooLarge(true), WithTemporary(true),
			WithUnauthorized(true),
		},
		fns:          allFuncs,
		expectations: []bool{true, true, true, true, true, true, true, true},
	},
	{
		opts: []OptsFunc{WithAccessDenied(true), WithTooLarge(true),
			WithConflict(true), WithNotFound(true), WithExists(true),
			WithRateLimit(true), WithTooLarge(false),
		},
		fns:          allFuncs,
		expectations: []bool{true, true, true, true, true, false, false, false},
	},
}
var allFuncs = []CheckerFn{
	IsAccessDenied, IsConflict, IsExist, IsNotFound,
	IsRateLimit, IsTooLarge, IsTemporary, IsUnauthorized,
}

func allFuncsExcept(fns ...CheckerFn) []CheckerFn {
	results := make([]CheckerFn, 0, len(allFuncs)-len(fns))
	lookup := make(map[string]struct{}, len(fns))
	for idx := range fns {
		lookup[getFunctionName(fns[idx])] = struct{}{}
	}

	for _, fn := range allFuncs {
		if _, found := lookup[getFunctionName(fn)]; found {
			continue
		}

		results = append(results, fn)
	}

	return results
}

func getFunctionName(i interface{}) string {
	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}
