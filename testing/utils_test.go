package utils

import (
	"testing"
	"time"
)


func TestBeginningOfMonth(t *testing.T) {
	dt := time.Date(2018, 2, 34, 0, 0, 0, 0, time.Local)
	resDt := BeginningOfMonth(dt)
	fmt.Println(resDt)
	if resDt.Day() != 1 {
		t.Error(resDt)
	}
}
