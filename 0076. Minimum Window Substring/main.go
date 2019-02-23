package main

import "fmt"

// Sliding Window Approach.
func minWindow(s string, t string) string {
	have := [128]int{}
	need := [128]int{}

	fmt.Println(have, need)

	for i := range t {
		need[t[i]]++
	}

	fmt.Println(need)

	size, total := len(s), len(t)

	min := size + 1

	res := "" // for return

	// s [i : j + 1]은 창입니다.
	// count는 기존 t의 문자 수를 세는 데 사용됩니다.
	// count == total은 필요한 모든 문자가 수집되었음을 의미합니다

	for i, j, count := 0, 0, 0; j < size; j++ {
		if have[s[j]] < need[s[j]] {
			// 창에 누락 된 문자가 나타납니다.
			count++
		}

		have[s[j]]++

		// 창문에서 필요한 문자를 잃지 않도록하십시오.
		// 가능한 한 크게 만든다.
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}

		width := j - i + 1
		if count == total && min > width {
			min = width
			res = s[i : j+1]
		}
	}

	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}
