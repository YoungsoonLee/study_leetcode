package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var res = []int{}

	var target, window [26]int
	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	check := func(i int) {
		fmt.Println("4: ", target, window)
		if window == target {
			res = append(res, i)
		}
	}

	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		//check(0)
		if i == len(p)-1 {
			check(0)
		} else if len(p) <= i {
			fmt.Println("1: ", target, window)
			window[s[i-len(p)]-'a']--
			fmt.Println("2: ", target, window)
			fmt.Println("3: ", i, i-len(p)+1)
			check(i - len(p) + 1)
		}
	}

	fmt.Println(target, window, res)

	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	findAnagrams(s, p)
}
