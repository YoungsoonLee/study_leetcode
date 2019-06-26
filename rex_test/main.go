package main

import (
	"fmt"
	"regexp"
)

func main() {
	displayname := "44eirweuri"
	re, _ := regexp.Compile("[^A-Za-z0-9]+")
	fmt.Println(re.FindString(displayname))

	//fmt.Println(matched)
}
