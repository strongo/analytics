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
			want: &event{message: message{event: "event1", properties: map[string]any{}}, category: "category1", action: "action1"},
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
