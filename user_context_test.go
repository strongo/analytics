package analytics

import (
	"context"
	"testing"
)

func TestUserContext(t *testing.T) {
	uc := NewUserContext("user123").
		SetUserLanguage("en-US").
		SetUserAgent("unit-test-runner")

	ctx := context.Background()

	pageView := NewPageview("unit-test", "TestUserContext").
		SetTitle("Test User Context")

	uc.QueueMessage(ctx, pageView)
}
