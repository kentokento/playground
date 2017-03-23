package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	u := []string{"seter", "test", "tesdfda"}
	fmt.Println(u)
	p := make([]string, 8)
	p = make([]string, 3)
	fmt.Println(p)
	m := copy(p, u)
	fmt.Println(p)
	fmt.Println(u)
	fmt.Println(m)

	u = u[:2]
	fmt.Println(u)
}
