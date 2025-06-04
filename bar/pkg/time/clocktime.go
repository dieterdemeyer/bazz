package time

import (
	"fmt"
	"time"
)

type ClockTime struct {
	Hours, Minutes, Seconds int
	Valid                   bool
}

func (ct ClockTime) ToString() string {
	return fmt.Sprintf("%s:%s:%s",
		fmt.Sprintf("%02d", ct.Hours),
		fmt.Sprintf("%02d", ct.Minutes),
		fmt.Sprintf("%02d", ct.Seconds),
	)
}

func ClockTimesFromHour(hour int) ClockTime {
	return ClockTime{hour, 0, 0, true}
}

func ClockTimesFromHours(hours ...int) []ClockTime {
	result := make([]ClockTime, len(hours))
	for i, hour := range hours {
		result[i] = ClockTimesFromHour(hour)
	}
	return result
}

func (ct ClockTime) ToTime(date time.Time) time.Time {
	if ct.Valid {
		return time.Date(date.Year(), date.Month(), date.Day(), ct.Hours, ct.Minutes, ct.Seconds, 0, date.Location())
	} else {
		return time.Time{}
	}
}

// Before reports whether the clocktime ct is before other.
func (ct ClockTime) Before(other ClockTime) bool {
	return ct.SecondsSinceDayStart() < other.SecondsSinceDayStart()
}

// After reports whether the clocktime ct is after other.
func (ct ClockTime) After(other ClockTime) bool {
	return ct.SecondsSinceDayStart() > other.SecondsSinceDayStart()
}

func (ct ClockTime) Compare(other ClockTime) int {
	return ct.SecondsSinceDayStart() - other.SecondsSinceDayStart()
}

func (ct ClockTime) SecondsSinceDayStart() int {
	return ct.Hours*60*60 + ct.Minutes*60 + ct.Seconds
}

func (ct ClockTime) Add(hours, minutes, seconds int) ClockTime {
	ct.Hours += hours
	ct.Minutes += minutes
	ct.Seconds += seconds
	return ct.Normalize()
}

func (ct ClockTime) Normalize() ClockTime {
	remainder := ct.Seconds / 60
	ct.Seconds %= 60
	ct.Minutes += remainder
	remainder = ct.Minutes / 60
	ct.Minutes %= 60
	ct.Hours += remainder
	return ct
}

type ClockTimeRange struct {
	Start ClockTime
	End   ClockTime
}

func (t ClockTimeRange) Before(other ClockTimeRange) bool {
	return t.Start.Before(other.Start) || (t.Start == other.Start && t.End.Before(other.End))
}

func (t ClockTimeRange) After(other ClockTimeRange) bool {
	return t.Start.After(other.Start) || (t.Start == other.Start && t.End.After(other.End))
}

func (t ClockTimeRange) IsValid() bool {
	if !t.Start.Valid || !t.End.Valid {
		return false
	}
	return t.Start.SecondsSinceDayStart() <= t.End.SecondsSinceDayStart()
}
