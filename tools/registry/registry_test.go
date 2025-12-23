package registry

import (
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New()
	if err == nil {
		t.Fatal("expected error not to be nil")
	}

	msg := `Config File ".env" Not Found in`
	if !strings.Contains(err.Error(), msg) {
		t.Fatalf("expected error to contain message %s, got %v", msg, err)
	}
}

func TestNewWithConfig(t *testing.T) {
	r, err := NewWithConfig(Config{
		Path: "../../tests/config",
		Name: "test_config",
	})
	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}

	if r == nil {
		t.Fatal("expected registry not to be nil")
	}

	greeting := r.GetString("TEST_GREETING")
	if greeting != "hello" {
		t.Fatalf("expected greeting value to be hello, got %s", greeting)
	}
}

func TestNewWithConfigEmpty(t *testing.T) {
	r, err := NewWithConfig(Config{
		Path: "../../tests/config",
		Name: "test_empty",
	})
	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}

	if r == nil {
		t.Fatal("expected registry not to be nil")
	}

	greeting := "hi"
	greetingVar := r.GetString("TEST_GREETING")

	if greetingVar != "" {
		greeting = greetingVar
	}

	// Since the file load is empty, we expect the `greeting` value not to be changed.
	if greeting != "hi" {
		t.Fatalf("expected greeting value to be hi, got %s", greeting)
	}
}

func TestNewWithConfigDefault(t *testing.T) {
	_, err := NewWithConfig(Config{})
	if err == nil {
		t.Fatal("expected error not to be nil")
	}

	msg := `Config File ".env" Not Found in`
	if !strings.Contains(err.Error(), msg) {
		t.Fatalf("expected error to contain message %s, got %v", msg, err)
	}
}

func TestNewWithConfigPrefix(t *testing.T) {
	r, err := NewWithConfig(Config{
		Path:      "../../tests/config",
		Name:      "test_config",
		EnvPrefix: "TESTING",
	})
	if err != nil {
		t.Fatalf("expected error to be nil, got %v", err)
	}

	if r == nil {
		t.Fatal("expected registry not to be nil")
	}

	if err := os.Setenv("TESTING_TEST_GREETING", "foo"); err != nil {
		t.Fatal(err)
	}

	greeting := r.GetString("TEST_GREETING")

	// Since we set a prefix, that should overwrite any variable in a file.
	if greeting != "foo" {
		t.Fatalf("expected greeting value to be foo, got %s", greeting)
	}
}
