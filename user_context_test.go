package analytics

import (
	"context"
	"testing"
)

func TestUserContext(t *testing.T) {
	uc := NewUserContext("user123").
		SetUserLanguage("en-US").
		SetUserAgent("unit-test-runner")

	if uc.GetUserID() != "user123" {
		t.Errorf("expected user123, got %q", uc.GetUserID())
	}
	if uc.GetUserLanguage() != "en-US" {
		t.Errorf("expected en-US, got %q", uc.GetUserLanguage())
	}
	if uc.GetUserAgent() != "unit-test-runner" {
		t.Errorf("expected unit-test-runner, got %q", uc.GetUserAgent())
	}
	if err := uc.Validate(); err != nil {
		t.Errorf("expected nil error from Validate, got %v", err)
	}

	ctx := context.Background()

	pageView := NewPageview("telegram", "/bot/TestBot/some/path")

	// Message without assigned user context
	uc.QueueMessage(ctx, pageView)

	// Message with same assigned user context
	uc.QueueMessage(ctx, pageView)

	// Message with different user context
	otherUc := NewUserContext("other456")
	otherUc.QueueMessage(ctx, pageView)

	// Message with non-*userContext type
	mockMsg := &message{event: "test", user: mockUserContext{}}
	uc.QueueMessage(ctx, mockMsg)

	// Nil receiver panics
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on nil userContext receiver")
			}
		}()
		var nilUc *userContext
		nilUc.QueueMessage(ctx, pageView)
	}()

	// Nil message panics
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on nil message")
			}
		}()
		uc.QueueMessage(ctx, nil)
	}()
}

