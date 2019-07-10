package main

import (
	"fmt"
	"sort"
)

func findRepeatedDnaSequences(s string) []string {
	var res []string
	if len(s) <= 10 {
		return nil
	}

	str := ""
	// rec는 각 subString의 발생 수를 기록합니다.
	rec := make(map[string]int, len(s)-9)
	for i := 0; i+10 <= len(s); i++ {
		str = s[i : i+10]
		//fmt.Println("for: ", str)
		if v := rec[str]; v == 1 {
			res = append(res, str)
		}
		rec[str]++
	}

	fmt.Println("before: ", res)
	sort.Strings(res)
	fmt.Println("after: ", res)

	return res

	/*
		res := make([]string, 0)
		for i := 0; i < len(s)-10; i++ {
			subs := s[i : i+10]
			if len(subs) < 10 {
				break
			}

			fmt.Println("subs: ", subs)

			if strings.Contains(s, subs) {
				res = append(res, subs)
			}
		}

		fmt.Println(res)
		return res
	*/
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	findRepeatedDnaSequences(s)
}
