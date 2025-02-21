package timetrack

import (
	"testing"
	"time"

	"github.com/rafael1mc/timetrack/internal/timeprovider"
)

func TestTimeStack(t *testing.T) {
	mockTime := timeprovider.NewMockTime(time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	SetupTimeProvider(mockTime)

	t.Run("Basic timing", func(t *testing.T) {
		root := NewNode("root node")

		mockTime.Advance(100 * time.Millisecond)
		duration := root.Stop()

		if duration != 100*time.Millisecond {
			t.Errorf("Expected duration 100ms, got %v", duration)
		}
	})

	t.Run("Nested timing", func(t *testing.T) {
		root := NewNode("root node")
		mockTime.Advance(50 * time.Millisecond)

		child := root.Branch("child node")
		mockTime.Advance(25 * time.Millisecond)
		childDuration := child.Stop()

		mockTime.Advance(100 * time.Millisecond)
		rootDuration := root.Stop()

		if childDuration != 25*time.Millisecond {
			t.Errorf("Expected child duration 25ms, got %v", child.duration)
		}

		if rootDuration != 175*time.Millisecond {
			t.Errorf("Expected root duration 175ms, got %v", rootDuration)
		}
	})

	t.Run("Parent stop, stops children", func(t *testing.T) {
		root := NewNode("root node")
		mockTime.Advance(50 * time.Millisecond)

		child := root.Branch("child node")
		mockTime.Advance(25 * time.Millisecond)

		rootDuration := root.Stop()
		childDuration := child.Stop()

		if childDuration != 25*time.Millisecond {
			t.Errorf("Expected child duration 25ms, got %v", child.duration)
		}

		if rootDuration != 75*time.Millisecond {
			t.Errorf("Expected root duration 75ms, got %v", rootDuration)
		}
	})

	t.Run("Multiple stop don't affect", func(t *testing.T) {
		root := NewNode("root node")

		mockTime.Advance(50 * time.Millisecond)
		rootDuration := root.Stop()

		mockTime.Advance(25 * time.Millisecond)
		rootDuration = root.Stop()

		if rootDuration != 50*time.Millisecond {
			t.Errorf("Expected root duration 50ms, got %v", rootDuration)
		}
	})
}
