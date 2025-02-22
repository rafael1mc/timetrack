package timetrack

import (
	"sync"

	"github.com/rafael1mc/timetrack/internal/timeprovider"
)

var defaultTimeProvider TimeProvider
var defaultReporter TimeReporter
var mu sync.Mutex

// timeProvider is a workaround so we don't have to keep passing the provider on every call
func timeProvider() TimeProvider {
	mu.Lock()
	defer mu.Unlock()

	if defaultTimeProvider == nil {
		defaultTimeProvider = timeprovider.NewRealTime()
	}

	return defaultTimeProvider
}

// SetupTimeProvider is used to change how time is tracked
// Useful for unit tests
func SetupTimeProvider(tp TimeProvider) {
	mu.Lock()
	defer mu.Unlock()

	defaultTimeProvider = tp
}

func reporter() TimeReporter {
	mu.Lock()
	defer mu.Unlock()

	if defaultReporter == nil {
		defaultReporter = NewSimpleTimeReporter()
	}

	return defaultReporter
}
