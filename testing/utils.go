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


//SomoniToInt convert somoni to int
// return 0 means could not convert
// it is not validate is it zero
func SomoniToInt(amt string) int64 {
	amt = strings.Replace(amt, ",", ".", -1)
	amts := strings.Split(amt, ".")
	if len(amts) == 1 {
		s, _ := strconv.ParseInt(amt, 10, 64)
		return s * 100
	}
	if ln := len(amts[1]); ln == 1 {
		amts[1] = amts[1] + "0"
	} else if ln > 2 {
		amts[1] = amts[1][:2]
	} else if ln == 0 {
		amts[1] = amts[1] + "00"
	}
	amt = amts[0] + amts[1]
	s, _ := strconv.ParseInt(amt, 10, 64)
	return s
}
