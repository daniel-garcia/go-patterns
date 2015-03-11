package example

import (
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {

	s := New("foo")
	if s == nil {
		t.Fatal("expected non-nil example")
	}

	if s.Name() != "foo" {
		t.Fatalf("expected name to be foo")
	}

	done := make(chan struct{})
	go func() {
		s.Close()
		close(done)
	}()
	select {
	case <-time.After(time.Second):
		t.Fatalf("timeout waiting for close")
	case <-done:
	}
}
