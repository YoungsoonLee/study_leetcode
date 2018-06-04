package main

import "fmt"

func Solution(S string, K int) int {
	temp := make([]string, 0)

	for i := len(S) - 1; i >= 0; i -= K {
		if i >= K {
			fmt.Println(i)
			fmt.Println(S[i-K:])
			temp = append(temp, S[i-K:])
			S = S[:i-K]
		} else {
			fmt.Println(S)
			temp = append(temp, S)
		}
	}

	fmt.Println(temp, len(temp))
	return len(temp)
}

func main() {
	s := "SMS messages are really short"
	k := 12
	Solution(s, k)
}
