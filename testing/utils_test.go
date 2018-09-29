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
	res := SomoniToInt("25,36,58")
	if res != 2536 {
		t.Fatalf("%v", res)
	}
}

func BenchmarkSomoniToInt(b *testing.B) {
	b.ResetTimer()
	str := "125458.3256.36"
	for n := 0; n < b.N; n++ {
		SomoniToInt(str)
	}
}

