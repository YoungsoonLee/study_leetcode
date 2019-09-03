package main

import (
	"fmt"
)

type data struct {
	index, count int
}

/*
func checkRecord(s string) bool {
	return !(strings.Count(s, "A") > 1 || strings.Contains(s, "LLL"))
}
*/

func checkRecord(s string) bool {
	acnt, lcnt := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			acnt++
			if acnt > 1 {
				return false
			}
		} else if s[i] == 'L' {
			lcnt++
			if lcnt > 2 {
				return false
			}
			continue // !!! key point
		}
		lcnt = 0 // !!!! key point. make not in a row. 도달 하지 않네. 위 continue
	}
	return true

}

func main() {
	a := "LLPALPs"
	fmt.Println(checkRecord(a))
}
