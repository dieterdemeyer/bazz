package time

import "time"

// Clock is an interface which can be used to get the current time.
// This makes it possible to mock times in test
type Clock interface {
	Now() time.Time
}

type systemClock struct{}

func (s systemClock) Now() time.Time {
	return time.Now()
}

// SystemClock is the global instance of the system clock.
var SystemClock Clock = systemClock{}

type alwaysSameTimeClock struct {
	usedTime time.Time
}

func (a alwaysSameTimeClock) Now() time.Time {
	return a.usedTime
}

// AlwaysSameTimeClock returns a clock that always returns the same time
func AlwaysSameTimeClock(t time.Time) Clock {
	return alwaysSameTimeClock{usedTime: t}
}
