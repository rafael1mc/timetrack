package timetrack

import (
	"time"

	"github.com/rafael1mc/timetrack/internal/timeprovider"
)

type TimeProvider interface {
	Now() time.Time
}

var timeProvider TimeProvider

// initTimeProvider is a workaround so we don't have to keep passing the provider on every call
func initTimeProvider() {
	if timeProvider == nil {
		SetupTimeProvider(timeprovider.NewRealTime())
	}
}

func SetupTimeProvider(tp TimeProvider) {
	timeProvider = tp
}
