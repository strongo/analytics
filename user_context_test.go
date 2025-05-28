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

	pageView := NewPageview("telegram", "/bot/TestBot/some/path").
		SetTitle("Test User Context")

	props := pageView.Properties()
	const expectedUrl = "tg://TestBot/some/path"
	if currentUrl := props["$current_url"]; currentUrl != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, currentUrl)
	}
	uc.QueueMessage(ctx, pageView)
}
