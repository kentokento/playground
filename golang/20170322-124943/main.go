package main

import (
	"fmt"
)

type behaviorChecker struct {
	checked map[string]checkedObj
}

type checkedObj struct {
	campaignID string
	isShow     bool
}

func (c checkedObj) Get() (string, bool) {
	return c.campaignID, c.isShow
}

func main() {
	fmt.Println("Hello, playground")
	m := map[string]int{}
	//m := make(map[string]int)

	m["apple"] = 150
	m["banana"] = 200
	m["lemon"] = 300

	fmt.Println(m["apple"])

	b := behaviorChecker{}
	b.checked = map[string]checkedObj{}
	fmt.Println(b.checked)
	id := "pickup"
	b.checked[id] = checkedObj{
		campaignID: "",
		isShow:     true,
	}
	fmt.Println(b.checked)

	//b := map[string]checkedObj{}
	//fmt.Println(b)
	//id := "pickup"
	//b[id] = checkedObj{
	//	campaignID: "",
	//	isShow:     true,
	//}
	//fmt.Println(b)
}
