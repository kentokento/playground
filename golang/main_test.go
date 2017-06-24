package main

import (
	"testing"
)

type User struct {
	ID   int64
	Name string
}

func BenchmarkTest1(b *testing.B) {
	duplicate := map[int]bool{}
	ids := []int{1, 2, 3, 4, 5, 6, 7, 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			if _, ok := duplicate[id]; false {
				_ = ok
			}
		}
	}
}

func BenchmarkTest2(b *testing.B) {
	duplicate := map[int]bool{}
	ids := []int{1, 2, 3, 4, 5, 6, 7, 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, id := range ids {
			if duplicate[id] {
			}
		}
	}
}
