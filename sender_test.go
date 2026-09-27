package analytics

import (
	"context"
	"testing"
)

type mockSender struct {
	sent []Message
}

func (m *mockSender) QueueMessage(ctx context.Context, msg Message) {
	m.sent = append(m.sent, msg)
}

func TestSender(t *testing.T) {
	ctx := context.Background()

	// Panics on nil ctx
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on nil ctx")
			}
		}()
		QueueMessage(nil, &message{})
	}()

	// Panics on nil msg
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic on nil msg")
			}
		}()
		QueueMessage(ctx, nil)
	}()

	// Invalid message
	invalidMsg := &message{event: ""}
	QueueMessage(ctx, invalidMsg)

	// Valid message with registered sender
	ms := &mockSender{}
	AddSender(ms)

	validMsg := &message{event: "test_event"}
	QueueMessage(ctx, validMsg)

	if len(ms.sent) == 0 {
		t.Error("expected mockSender to receive message")
	}
}
