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

type message struct {
	ApiClientID string
	event       string
	category    string
	user        UserContext
}

func (m *message) Validate() error {
	if m.event == "" {
		return errors.New("event is required")
	}
	if m.ApiClientID == "" {
		return errors.New("missing api client id")
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
