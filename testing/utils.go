package utils

import (
	"time"
)

// BeginningOfMonth beginning of month
func BeginningOfMonth(dt time.Time) time.Time {
	y, m, _ := dt.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

// EndOfMonth end of month
func EndOfMonth(dt time.Time) time.Time {
	return BeginningOfMonth(dt).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginningOfDay beginning of month
func BeginningOfDay(dt time.Time) time.Time {
	y, m, d := dt.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}


