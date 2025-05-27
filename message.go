package analytics

import "errors"

type Message interface {
	Validate() error
	GetApiClientID() string
	User() UserContext
	SetUserContext(user UserContext)
}

type message struct {
	ApiClientID string
	user        UserContext
}

func (m *message) GetApiClientID() string {
	return m.ApiClientID
}

func (m *message) User() UserContext {
	return m.user
}

func (m *message) SetUserContext(user UserContext) {
	m.user = user
}

func (m *message) Validate() error {
	if m.ApiClientID == "" {
		return errors.New("missing api client id")
	}
	if err := m.user.Validate(); err != nil {
		return err
	}
	return nil
}
