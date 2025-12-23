package sentinel

import "testing"

func TestNew(t *testing.T) {
	app := New()

	if app == nil {
		t.Fatal("expected initialized Sentinel instance, got nil")
	}

	if app.App == nil {
		t.Fatal("expected app.App to be initialized, got nil")
	}
}

func TestNewWithConfig(t *testing.T) {
	app := NewWithConfig(Config{})

	if app == nil {
		t.Fatal("expected initialized Sentinel instance, got nil")
	}

	if app.App == nil {
		t.Fatal("expected app.App to be initialized, got nil")
	}
}
