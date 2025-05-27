package analytics

import (
	"context"
	"github.com/strongo/logus"
)

type Sender interface {
	QueueMessage(message Message)
}

var senders []Sender

func QueueMessage(ctx context.Context, msg Message) {
	if err := msg.Validate(); err != nil {
		logus.Errorf(ctx, "queued invalid message to analytics: %s", err)
		return
	}
	for _, sender := range senders {
		sender.QueueMessage(msg)
	}
}

func AddSender(sender Sender) {
	senders = append(senders, sender)
}
