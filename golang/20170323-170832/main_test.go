package main_test

import (
	"fmt"
	"strings"
	"testing"
)

var (
	max = 100
)

func initList() map[int64]bool {
	l := map[int64]bool{}
	for i := 1000; i < 30000; i++ {
		l[int64(i)] = true
	}
	return l
}

func BenchmarkTest1(b *testing.B) {
	partnerCommunityIDs := initList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idString := ""
		cnt := 0
		for id := range partnerCommunityIDs {
			idString += fmt.Sprintf("%d,", id)
			if cnt >= max {
				break
			}
			cnt++
		}
	}
}

func BenchmarkTest2(b *testing.B) {
	partnerCommunityIDs := initList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idString := ""
		cnt := 0
		var m2 []string
		for id := range partnerCommunityIDs {
			m2 = append(m2, fmt.Sprintf("%d", id))
			if cnt >= max {
				break
			}
			cnt++
		}
		idString = strings.Trim(strings.Replace(fmt.Sprint(m2), " ", ",", -1), "[]")
		_ = idString
	}
}

func BenchmarkTest3(b *testing.B) {
	partnerCommunityIDs := initList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idString := ""
		cnt := 0
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
		_ = idString
	}
}

func BenchmarkTest4(b *testing.B) {
	partnerCommunityIDs := initList()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idString := ""
		cnt := 0
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
		_ = idString
	}
}
