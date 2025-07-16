package analytics

import (
	"context"
	"github.com/strongo/logus"
)

type Sender interface {
	QueueMessage(ctx context.Context, message Message)
}

var senders []Sender

func QueueMessage(ctx context.Context, msg Message) {
	if ctx == nil {
		panic("analytics.QueueMessage(ctx=nil)")
	}
	if msg == nil {
		panic("analytics.QueueMessage(msg=nil)")
	}
	if err := msg.Validate(); err != nil {
		logus.Errorf(ctx, "analytics.QueueMessage() receieved an invalid message, validation error %T: %s", err, err)
		return
	}
	for _, sender := range senders {
		sender.QueueMessage(ctx, msg)
	}
}

func AddSender(sender Sender) {
	senders = append(senders, sender)
}
