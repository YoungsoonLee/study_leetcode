package main

import "fmt"

func findTheDifference(s string, t string) byte {
	rec := make([]int, 26)
	for i := range s {
		rec[s[i]-'a']--
		rec[t[i]-'a']++
	}

	rec[t[len(t)-1]-'a']++ // !!!

	var i int
	for i = 0; i < 26; i++ {
		if rec[i] == 1 {
			break
		}
	}

	return byte('a' + i)
}

func main() {
	s := "abcdX"
	t := "abcdXe"
	fmt.Println(string(findTheDifference(s, t)))
}
