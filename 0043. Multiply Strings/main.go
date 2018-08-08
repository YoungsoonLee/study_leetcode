package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// string is converted to []byte, it is easy to get the specific value on the corresponding bit
	bsi := []byte(num1)
	bsj := []byte(num2)

	fmt.Println("hsi: ", bsi)
	fmt.Println("hsj: ", bsj)

	// The length of temp can only be len(num1)+len(num2) or len(num1)+len(num2)-1
	// Choose long, so that you don’t have enough
	temp := make([]int, len(num1)+len(num2))
	for i := 0; i < len(bsi); i++ {
		for j := 0; j < len(bsj); j++ {
			// The product on each bit can be directly stored in the corresponding position in temp
			temp[i+j+1] += int(bsi[i]-'0') * int(bsj[j]-'0')
		}
	}

	fmt.Println("temp: ", temp)

	// Unified processing carry
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	fmt.Println("temp: ", temp)

	// When num1 and num2 are small, the first position of temp is 0.
	// To avoid output results starting with 0, you need to remove the 0 first place of temp
	if temp[0] == 0 {
		temp = temp[1:]
	}

	// conversion result
	// temp is chosen to be []int instead of []byte because
	// In Go, the base structure of byte is uint8, and the maximum value is 255.
	// temp will overflow if you don’t consider the carry
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}
	fmt.Println("res: ", res)

	return string(res)
}

func main() {
	n1 := "12"
	n2 := "45"

	fmt.Println(multiply(n1, n2))

}
