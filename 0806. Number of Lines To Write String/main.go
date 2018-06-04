package main

import "fmt"

func numberOfLines(widths []int, S string) []int {
	res := []int{0, 0}
	if len(S) == 0 {
		return res
	}

	res[0] = 1

	for i := 0; i < len(S); i++ {
		fmt.Println(res)
		if res[1]+widths[S[i]-'a'] > 100 {
			res[0]++
			res[1] = widths[S[i]-'a']
		} else {
			res[1] += widths[S[i]-'a']
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	w := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"

	numberOfLines(w, s)
}
