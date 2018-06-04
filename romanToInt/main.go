package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanSymbolValue := make(map[string]int)
	romanSymbolValue["I"] = 1
	romanSymbolValue["V"] = 5
	romanSymbolValue["X"] = 10
	romanSymbolValue["L"] = 50
	romanSymbolValue["C"] = 100
	romanSymbolValue["D"] = 500
	romanSymbolValue["M"] = 1000

	result := 0
	for i := 0; i < len(s)-1; i++ {
		//fmt.Println(result)
		if romanSymbolValue[string(s[i])] < romanSymbolValue[string(s[i+1])] {
			result -= romanSymbolValue[string(s[i])]
		} else {
			result += romanSymbolValue[string(s[i])]
		}
	}

	result += romanSymbolValue[string(s[len(s)-1])]

	if result < 1 || result > 3999 {
		return 0 //, errors.New("over input range. input range is from 0 to 3999")
	}
	return result

}

func main() {
	fmt.Println(romanToInt("MCDLXXVI"))
}
