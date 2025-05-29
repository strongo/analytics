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

	pageView := NewPageview("telegram", "/bot/TestBot/some/path")

	uc.QueueMessage(ctx, pageView)
}
