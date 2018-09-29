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

func TestSomoniToInt(t *testing.T) {
	for k, tt := range []struct {
		inp    string
		Result int64
	}{
		{inp: "1254", Result: 125400},
		{inp: "1,254", Result: 125},
		{inp: "12.54", Result: 1254},
		{inp: "125,4", Result: 12540},
		{inp: "1254,", Result: 125400},
		{inp: "1254.", Result: 125400},
		{inp: "1254.0", Result: 125400},
		{inp: "0.01", Result: 1},
		{inp: "01.03", Result: 103},
		{inp: "-5.03", Result: -503},
		{inp: "-5.3", Result: -530},
	} {
		if res := SomoniToInt(tt.inp); res != tt.Result {
			t.Fatalf("%d - expected diram of %v to be %d, got %d\n", k, tt.inp, tt.Result, res)
		}
	}
}

func BenchmarkSomoniToInt(b *testing.B) {
	b.ResetTimer()
	str := "125458.3256.36"
	for n := 0; n < b.N; n++ {
		SomoniToInt(str)
	}
}

