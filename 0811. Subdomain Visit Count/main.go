package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains))

	for _, domain := range cpdomains {
		d, n := parse(domain) // d: discuss.leetcode.com , n: 9001
		isNew := true
		for isNew {
			m[d] += n
			d, isNew = cut(d)
		}
	}
	fmt.Println("m: ", m)
	return getResult(m)
}

func parse(s string) (string, int) {
	ss := strings.Split(s, " ")
	n, _ := strconv.Atoi(ss[0])
	fmt.Println("parse: ", ss, n)
	return ss[1], n
}

func cut(s string) (string, bool) {
	idx := strings.Index(s, ".")
	if idx == -1 {
		return "", false
	}
	fmt.Println("cut: ", s, idx, s[idx+1:])
	return s[idx+1:], true
}

func getResult(m map[string]int) []string {
	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}

func main() {
	//a := []string{"9001 discuss.leetcode.com"}
	a := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	subdomainVisits(a)
}
