package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	// Primary
	ID   string `json:"id"`
	Name string `json:"name"`

	// Sub
	Birthday           string `json:"birthday"`
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	Gender             string `json:"gender"`
	LastName           string `json:"last_name"`
	Link               string `json:"link"`
	Locale             string `json:"locale"`
	Location           string `json:"location"`
	RelationshipStatus string `json:"relationship_status"`
	Verified           bool   `json:"verified"`
}

func testStr() string {
	return "testeste"
}

const test = testStr()

func main() {
	fmt.Println("Hello, playground")
	body, _ := json.Marshal(User{})
	fmt.Println(string(body))
	fmt.Println(test)
}
