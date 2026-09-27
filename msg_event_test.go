package analytics

import (
	"reflect"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
		name     string
		category string
		action   string
	}
	tests := []struct {
		name string
		args args
		want *event
	}{
		{
			name: "should_pass",
			args: args{
				name:     "event1",
				category: "category1",
				action:   "action1",
			},
			want: &event{message: newMessage("event1"), category: "category1", action: "action1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvent(tt.args.name, tt.args.category, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventMethodsAndValidation(t *testing.T) {
	ev := NewEvent("name1", "cat1", "act1")

	ev.SetTitle("Title1").
		SetAction("act2").
		SetLabel("lbl1").
		SetValue(42)

	if ev.Title() != "Title1" {
		t.Errorf("expected Title1, got %q", ev.Title())
	}
	if ev.Action() != "act2" {
		t.Errorf("expected act2, got %q", ev.Action())
	}
	if ev.Label() != "lbl1" {
		t.Errorf("expected lbl1, got %q", ev.Label())
	}
	if ev.Value() != 42 {
		t.Errorf("expected 42, got %d", ev.Value())
	}

	if err := ev.Validate(); err != nil {
		t.Errorf("expected valid event, got %v", err)
	}

	// Message validation failure (empty event name)
	invalidMsg := &event{message: message{event: ""}, category: "cat", action: "act"}
	if err := invalidMsg.Validate(); err == nil {
		t.Error("expected error for empty event name")
	}

	// Missing category
	noCat := &event{message: newMessage("ev"), category: "", action: "act"}
	if err := noCat.Validate(); err == nil {
		t.Error("expected error for empty category")
	}

	// Missing action
	noAct := &event{message: newMessage("ev"), category: "cat", action: ""}
	if err := noAct.Validate(); err == nil {
		t.Error("expected error for empty action")
	}
}

