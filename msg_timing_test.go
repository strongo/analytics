package analytics

import (
	"testing"
	"time"
)

func TestTiming(t *testing.T) {
	dur := 150 * time.Millisecond
	timing := NewTiming("query_time", dur)

	if timing.Duration() != dur {
		t.Errorf("expected %v, got %v", dur, timing.Duration())
	}

	if err := timing.Validate(); err != nil {
		t.Errorf("expected valid timing, got %v", err)
	}

	zeroTiming := NewTiming("zero_time", 0)
	if err := zeroTiming.Validate(); err == nil {
		t.Error("expected validation error for 0 duration")
	}
}
