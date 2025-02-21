package timeprovider

import "time"

type MockTime struct {
	currentTime time.Time
}

func NewMockTime(initTime time.Time) *MockTime {
	return &MockTime{currentTime: initTime}
}

func (m *MockTime) Now() time.Time {
	return m.currentTime
}

func (m *MockTime) Advance(d time.Duration) {
	m.currentTime = m.currentTime.Add(d)
}
