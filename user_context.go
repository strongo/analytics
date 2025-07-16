package analytics

import (
	"context"
	"github.com/strongo/logus"
)

type UserContext interface {
	GetUserID() string

	Validate() error

	GetUserLanguage() string
	SetUserLanguage(language string) UserContext

	GetUserAgent() string
	SetUserAgent(userAgent string) UserContext

	QueueMessage(ctx context.Context, msg Message)
}

var _ UserContext = (*userContext)(nil)

type userContext struct {
	UserID       string `key:"uid" json:"uid"`
	UserLanguage string `key:"ul" json:"ul"`
	UserAgent    string `key:"ua" json:"ua"`
}

func (v *userContext) GetUserID() string {
	return v.UserID
}

func (v *userContext) GetUserLanguage() string {
	return v.UserLanguage
}

func (v *userContext) GetUserAgent() string {
	return v.UserAgent
}

func (v *userContext) Validate() error {
	return nil
}

func NewUserContext(userID string) UserContext {
	return &userContext{UserID: userID}
}

func (v *userContext) SetUserLanguage(s string) UserContext {
	v.UserLanguage = s
	return v
}

func (v *userContext) SetUserAgent(s string) UserContext {
	v.UserAgent = s
	return v
}

func (v *userContext) QueueMessage(ctx context.Context, msg Message) {
	if v == nil {
		panic("QueueMessage() called on nil userContext")
	}
	if msg == nil {
		panic("func *userContext.QueueMessage(msg=nil)")
	}
	if user := msg.User(); user != nil {
		if u, ok := user.(*userContext); !ok || u != v {
			logus.Errorf(ctx, "an attemp to send an analytics message with an assigned user context through a different user")
			return
		}
	}
	msg.SetUserContext(v)
	QueueMessage(ctx, msg)
}
