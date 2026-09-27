package analytics

import (
	"errors"
	"testing"
)

type mockUserContext struct {
	UserContext
	validErr error
}

func (m mockUserContext) Validate() error {
	return m.validErr
}

func TestMessage(t *testing.T) {
	m := newMessage("test_event")

	if m.Event() != "test_event" {
		t.Errorf("expected test_event, got %q", m.Event())
	}

	m.ApiClientID = "client_1"
	if m.GetApiClientID() != "client_1" {
		t.Errorf("expected client_1, got %q", m.GetApiClientID())
	}

	m.SetCategory("cat1")
	if m.Category() != "cat1" {
		t.Errorf("expected cat1, got %q", m.Category())
	}

	if m.Properties() == nil {
		t.Error("expected non-nil properties")
	}

	if err := m.Validate(); err != nil {
		t.Errorf("expected valid message without user, got %v", err)
	}

	// Empty event
	emptyMsg := newMessage("")
	if err := emptyMsg.Validate(); err == nil {
		t.Error("expected error for empty event")
	}

	// User validation error
	uErr := mockUserContext{validErr: errors.New("user invalid")}
	m.SetUserContext(uErr)
	if m.User() != uErr {
		t.Errorf("expected user context to be set")
	}
	if err := m.Validate(); err == nil || err.Error() != "user invalid" {
		t.Errorf("expected user invalid error, got %v", err)
	}

	// Valid user
	uValid := mockUserContext{validErr: nil}
	m.SetUserContext(uValid)
	if err := m.Validate(); err != nil {
		t.Errorf("expected valid message with valid user, got %v", err)
	}
}
