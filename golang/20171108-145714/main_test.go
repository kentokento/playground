package main

import (
	"testing"
	"time"
)

func BenchmarkTest1(b *testing.B) {

	week := map[time.Weekday]bool{
		time.Sunday:    false,
		time.Monday:    false,
		time.Tuesday:   false,
		time.Wednesday: false,
		time.Thursday:  false,
		time.Friday:    false,
		time.Saturday:  false,
	}
	tt := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range tt {
			_, _ = week[v]
		}
	}
}

func BenchmarkTest2(b *testing.B) {

	week := map[string]bool{
		time.Sunday.String():    false,
		time.Monday.String():    false,
		time.Tuesday.String():   false,
		time.Wednesday.String(): false,
		time.Thursday.String():  false,
		time.Friday.String():    false,
		time.Saturday.String():  false,
	}
	tt := []string{
		time.Sunday.String(),
		time.Monday.String(),
		time.Tuesday.String(),
		time.Wednesday.String(),
		time.Thursday.String(),
		time.Friday.String(),
		time.Saturday.String(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range tt {
			_, _ = week[v]
		}
	}
}
