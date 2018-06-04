package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	jMap := make(map[rune]bool, len(J))

	for _, j := range J {
		jMap[j] = true
	}
	fmt.Println(jMap)
	count := 0
	for _, c := range S {
		if _, exists := jMap[c]; exists {
			count++
		}
	}

	return count
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}
