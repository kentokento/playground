package main

import (
	"fmt"
)

type Hoge struct {
	Str string
}

func newHoge() *Hoge {
	return &Hoge{Str: ""}
}

func main() {
	fmt.Println("Hello, playground")

	a := newHoge()
	a.Str = "test"
	fmt.Println(a.Str)
	b := newHoge()
	b.Str = "test2"
	fmt.Println(a.Str)
	fmt.Println(b.Str)
}
