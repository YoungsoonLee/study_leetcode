package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	fmt.Println(err)
	return err == nil || strings.Contains(err.Error(), "value out of range")
}

func main() {
	args := os.Args
	fmt.Println(isNumber(args[1]))
}
