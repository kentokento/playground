package main

import (
	"fmt"
	"strings"
)

var (
	max = 1
)

func initList() map[int64]bool {
	l := map[int64]bool{}
	for i := 1000; i < (1000 + 10); i++ {
		l[int64(i)] = true
	}
	return l
}

func mainTest1() {
	partnerCommunityIDs := initList()
	fmt.Println(partnerCommunityIDs)
	idString := ""
	cnt := 0
	for id := range partnerCommunityIDs {
		idString += fmt.Sprintf("%d,", id)
		if cnt >= max {
			break
		}
		cnt++
	}
	fmt.Println(idString)
}

//func mainTest2() {
//	partnerCommunityIDs := initList()
//
//	idString := ""
//	a := partnerCommunityIDs[:max]
//	idString = strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
//	fmt.Println(idString)
//}
//
//func mainTest3() {
//	partnerCommunityIDs := initList()
//
//	idString := ""
//	a := partnerCommunityIDs[:max]
//	idString = strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ","), "[]")
//	fmt.Println(idString)
//}
//
//func mainTest4() {
//	partnerCommunityIDs := initList()
//
//	idString := ""
//	a := partnerCommunityIDs[:max]
//	idString = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
//	fmt.Println(idString)
//}

func mainTest5() {
	partnerCommunityIDs := initList()
	fmt.Println(partnerCommunityIDs)
	idString := ""
	cnt := 1
	var m2 []string
	for id := range partnerCommunityIDs {
		m2 = append(m2, fmt.Sprintf("%d,", id))
		if cnt >= max {
			break
		}
		cnt++
	}
	//idString = strings.Trim(strings.Replace(fmt.Sprint(m2), " ", ",", -1), "[]")
	idString = strings.Trim(fmt.Sprint(m2), "[]")
	fmt.Println(idString)
}

func mainTest6() {
	partnerCommunityIDs := initList()
	fmt.Println(partnerCommunityIDs)
	idString := ""
	cnt := 1
	var m2 []byte
	for id := range partnerCommunityIDs {
		m2 = append(m2, fmt.Sprintf("%d,", id)...)
		if cnt >= max {
			break
		}
		cnt++
	}
	//idString = strings.Trim(strings.Replace(fmt.Sprint(m2), " ", ",", -1), "[]")
	//idString = strings.Trim(fmt.Sprint(m2), "[]")
	idString = string(m2)
	fmt.Println(idString)
}

func main() {
	mainTest1()
	mainTest1()
	//	mainTest2()
	//	mainTest3()
	//	mainTest4()
	//mainTest5()
	mainTest6()
	mainTest6()
}
