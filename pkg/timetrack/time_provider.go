package timetrack

import (
	"time"
)

type TimeProvider interface {
	Now() time.Time
}
