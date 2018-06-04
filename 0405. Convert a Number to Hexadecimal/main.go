package main

import "fmt"

var h = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hex := ""
	for i := 0; i < 8 && num != 0; i++ { // 왜 8 일까???
		fmt.Println("1: ", num&15, hex)
		hex = h[num&15] + hex
		fmt.Println("2: ", hex, num)
		num >>= 4 // 왜 >> 4 일까??? 나누기 니깐... >>,
		fmt.Println(num)
	}
	fmt.Println(hex)
	return hex
}

func main() {
	a := 9
	toHex(a)
}
