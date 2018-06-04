package main

import (
	"fmt"
	"strings"
)

type data struct {
	index, count int
}

func checkRecord(s string) bool {
	return !(strings.Count(s, "A") > 1 || strings.Contains(s, "LLL"))
}

func main() {
	a := "LALL"
	fmt.Println(checkRecord(a))
}
