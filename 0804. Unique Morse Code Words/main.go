package main

import (
	"fmt"
	"strings"
)

var table = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	res := make(map[string]bool, len(words))
	for _, w := range words {
		var b strings.Builder // go 1.10
		for i := 0; i < len(w); i++ {
			b.WriteString(table[w[i]-'a'])
		}
		res[b.String()] = true
	}
	fmt.Println(res, len(res))
	return len(res)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	uniqueMorseRepresentations(words)
}
