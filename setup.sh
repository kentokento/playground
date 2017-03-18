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
