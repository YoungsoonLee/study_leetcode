package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	args := os.Args
	min, _ := strconv.ParseFloat(args[1], 64)
	max, _ := strconv.ParseFloat(args[1], 64)

	for i := 2; i < len(args); i++ {
		n, _ := strconv.ParseFloat(args[1], 64)
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	//checkingss
	//
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
