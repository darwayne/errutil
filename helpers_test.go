package errutil

import "testing"

func TestAccessDenied(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewAccessDenied("hey")
		assertTrue(t, IsAccessDenied(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewAccessDenied("hey")
		err = WrapAccessDenied(err)
		assertTrue(t, IsAccessDenied(err))
		t.Logf("%+v", err)
	})
}

func TestExists(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewExists("hey")
		assertTrue(t, IsExist(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewExists("hey")
		err = WrapExists(err)
		assertTrue(t, IsExist(err))
		t.Logf("%+v", err)
	})
}

func TestConflict(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewConflict("hey")
		assertTrue(t, IsConflict(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewConflict("hey")
		err = WrapConflict(err)
		assertTrue(t, IsConflict(err))
		t.Logf("%+v", err)
	})
}

func TestNotFound(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewNotFound("hey")
		assertTrue(t, IsNotFound(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewNotFound("hey")
		err = WrapNotFound(err)
		assertTrue(t, IsNotFound(err))
		t.Logf("%+v", err)
	})
}

func TestRateLimit(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewRateLimit("hey")
		assertTrue(t, IsRateLimit(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewRateLimit("hey")
		err = WrapRateLimit(err)
		assertTrue(t, IsRateLimit(err))
		t.Logf("%+v", err)
	})
}

func TestTemporary(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewTemporary("hey")
		assertTrue(t, IsTemporary(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewTemporary("hey")
		err = WrapTemporary(err)
		assertTrue(t, IsTemporary(err))
		t.Logf("%+v", err)
	})
}

func TestTooLarge(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewTooLarge("hey")
		assertTrue(t, IsTooLarge(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewTooLarge("hey")
		err = WrapTooLarge(err)
		assertTrue(t, IsTooLarge(err))
		t.Logf("%+v", err)
	})
}

func TestTooManyRequests(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewTooManyRequests("hey")
		assertTrue(t, IsTooManyRequests(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewTooManyRequests("hey")
		err = WrapTooManyRequests(err)
		assertTrue(t, IsTooManyRequests(err))
		t.Logf("%+v", err)
	})
}

func TestUnauthorized(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		err := NewUnauthorized("hey")
		assertTrue(t, IsUnauthorized(err))
		t.Logf("%+v", err)
	})

	t.Run("Wrap", func(t *testing.T) {
		err := NewUnauthorized("hey")
		err = WrapUnauthorized(err)
		assertTrue(t, IsUnauthorized(err))
		t.Logf("%+v", err)
	})
}

func assertTrue(t *testing.T, val bool) {
	t.Helper()
	if !val {
		t.Fatal("expected true")
	}
}
