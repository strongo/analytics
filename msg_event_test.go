package analytics

import (
	"reflect"
	"testing"
)

func TestNewEvent(t *testing.T) {
	type args struct {
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
				category: "test1",
				action:   "test2",
			},
			want: &event{category: "test1", action: "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEvent(tt.args.category, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
