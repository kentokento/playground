package main

import (
	"testing"
)

type User struct {
	ID   int64
	Name string
}

func initList() []User {
	partners := []User{
		{ID: 10001, Name: "test1"},
		{ID: 10001, Name: "test1"},
		{ID: 10001, Name: "test1"},
		{ID: 10001, Name: "test1"},
		{ID: 10001, Name: "test1"},
	}
	return partners
}

func BenchmarkAppend(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var exportPartners []User
		partners := initList()
		exportPartners = append(exportPartners, partners...)
	}
}

func BenchmarkCopy(b *testing.B) {
	//exportPartners := make([]User, 4)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		partners := initList()
		tmp := partners[:4]
		partners = tmp
	}
}
