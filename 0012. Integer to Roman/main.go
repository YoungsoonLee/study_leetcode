package main

import "fmt"

func intToRoman(num int) string {
	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}

	fmt.Println(num/1000, num/100%10, num/10%10, d[0][num%10])
	return d[3][num/1000] + d[2][num/100%10] + d[1][num/10%10] + d[0][num%10]
}

func main() {
	a := 800
	intToRoman(a)
}
