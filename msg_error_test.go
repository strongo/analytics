package analytics

import (
	"errors"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	err := errors.New("test error")
	msg := NewErrorMessage(err)

	if msg.ErrorText() != "test error" {
		t.Errorf("expected 'test error', got %q", msg.ErrorText())
	}

	if err := msg.Validate(); err == nil || err.Error() != "test error" {
		t.Errorf("expected validation to return err, got %v", err)
	}

	nilMsg := NewErrorMessage(nil)
	if err := nilMsg.Validate(); err != nil {
		t.Errorf("expected nil validation error, got %v", err)
	}
}
