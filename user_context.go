package analytics

import "context"

type UserContext struct {
	UserID       string `key:"uid" json:"uid"`
	UserLanguage string `key:"ul" json:"ul"`
	UserAgent    string `key:"ua" json:"ua"`
}

func (v *UserContext) Validate() error {
	return nil
}

func NewUserContext(userID string) *UserContext {
	return &UserContext{UserID: userID}
}

func (v *UserContext) SetUserLanguage(s string) *UserContext {
	v.UserLanguage = s
	return v
}

func (v *UserContext) SetUserAgent(s string) *UserContext {
	v.UserAgent = s
	return v
}

func (v *UserContext) QueueMessage(ctx context.Context, msg Message) {
	msg.SetUserContext(*v)
	QueueMessage(ctx, msg)
}
