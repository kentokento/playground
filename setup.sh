#!/bin/sh

newGolang() {
echo 'package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}' >> ./main.go
}

newBench() {
echo 'package main

import (
	"testing"
)

type User struct {
	ID   int64
	Name string
}

func BenchmarkTest1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
	}
}' >> ./main_test.go
}
