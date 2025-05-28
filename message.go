package analytics

import "errors"

type Message interface {
	Event() string
	Validate() error
	GetApiClientID() string
	User() UserContext
	SetUserContext(user UserContext)
	Category() string
	SetCategory(category string)
}

var _ Message = (*message)(nil)

func newMessage(event string) message {
	return message{event: event, properties: make(Properties, 10)}
}

type message struct {
	ApiClientID string
	event       string
	category    string
	user        UserContext
	properties  Properties
}

func (m *message) Properties() Properties {
	return m.properties
}

func (m *message) Validate() error {
	if m.event == "" {
		return errors.New("event is required")
	}
	if err := m.user.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *message) GetApiClientID() string {
	return m.ApiClientID
}

func (m *message) Event() string {
	return m.event
}

func (m *message) Category() string {
	return m.category
}

func (m *message) SetCategory(category string) {
	m.category = category
}

func (m *message) User() UserContext {
	return m.user
}

func (m *message) SetUserContext(user UserContext) {
	m.user = user
}
