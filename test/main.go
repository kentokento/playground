package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, playground")
	// imageURL := "https://graph.facebook.com/1429250183869006/picture?type=large"
	// imageURL := "https://graph.facebook.com/103670800481589/picture?type=large"
	imageURL := "https://graph.facebook.com/1705340493033069/picture?type=large"
	resp, err := http.DefaultClient.Get(imageURL)
	fmt.Println(err)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("error")
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Location"))
	fmt.Println(resp.Header.Get("Content-Disposition"))
}
