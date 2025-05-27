package analytics

import (
	"fmt"
	"time"
)

var _ Message = (*Timing)(nil)

type Timing struct {
	message
	Duration time.Duration `json:"duration"`
}

func (timing *Timing) Validate() error {
	if timing.Duration == 0 {
		return fmt.Errorf("timing duration is zero")
	}
	return nil
}
