package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func binaryStringToInt(s string) int {
	rs := 0
	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(s[i]))
		rs += v * int(math.Pow(float64(2), float64(j)))
		j++
	}
	fmt.Println("binaryStringToInt: ", rs)
	return rs
}

func intToBinaryString(i int) string {
	bs := fmt.Sprintf("%b", i)
	return bs
}

func addBinary(a string, b string) string {
	rs := binaryStringToInt(a) + binaryStringToInt(b)
	return intToBinaryString(rs)
}

// add bit is XOR
func addBinary2(a string, b string) string {
	rs := make([]string, 0)
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 {
		sum := 0

		if i >= 0 {
			v1, _ := strconv.Atoi(string(a[i]))
			sum += v1
			i--
		}

		if j >= 0 {
			v2, _ := strconv.Atoi(string(b[j]))
			sum += v2
			j--
		}
		sum += carry

		//fmt.Println("loop sum: ", sum)

		carry = sum / 2 // !!!!
		rs = append(rs, strconv.Itoa(sum%2))

		//fmt.Println("loop rs: ", rs)
	}

	if carry != 0 {
		rs = append(rs, strconv.Itoa(carry))
	}

	//fmt.Println("rs: ", reverse(rs))
	return strings.Join(reverse(rs), "")
}

func reverse(a []string) []string {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary2(a, b))
}
