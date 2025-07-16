package analytics

import "errors"

var _ Message = (*event)(nil)
var _ Event = (*event)(nil)

type Event interface {
	Message

	Action() string
	SetAction(string) Event

	Label() string
	SetLabel(string) Event

	Value() uint
	SetValue(uint) Event
}

func NewEvent(name, category, action string) Event {
	return &event{message: newMessage(name), category: category, action: action}
}

type event struct {
	message
	category string `key:"ec" required:"true"`
	action   string `key:"ea" required:"true"`
	label    string `key:"el"`
	value    uint   `key:"ev"`
}

func (v *event) Action() string {
	return v.action
}

func (v *event) Label() string {
	return v.label
}

func (v *event) Value() uint {
	return v.value
}

func (v *event) Validate() error {
	if err := v.message.Validate(); err != nil {
		return err
	}
	if v.category == "" {
		return errors.New("category is required")
	}
	if v.action == "" {
		return errors.New("action is required")
	}
	return nil
}

func (v *event) SetAction(action string) Event {
	v.action = action
	return v
}

func (v *event) SetLabel(label string) Event {
	v.label = label
	return v
}

func (v *event) SetValue(value uint) Event {
	v.value = value
	return v
}
