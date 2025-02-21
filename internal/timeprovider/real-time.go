package timeprovider

import "time"

type RealTime struct{}

func NewRealTime() *RealTime {
	return &RealTime{}
}

func (r RealTime) Now() time.Time {
	return time.Now()
}
