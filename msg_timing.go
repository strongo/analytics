package analytics

import (
	"fmt"
	"time"
)

var _ Message = (*timing)(nil)

type Timing interface {
	Message
	Duration() time.Duration
}

func NewTiming(event string, duration time.Duration) Timing {
	return &timing{message: newMessage(event), duration: duration}
}

type timing struct {
	message
	duration time.Duration `key:"td" required:"true"`
}

func (t *timing) Duration() time.Duration {
	return t.duration
}

func (t *timing) Validate() error {
	if t.duration == 0 {
		return fmt.Errorf("timing duration is zero")
	}
	return nil
}
