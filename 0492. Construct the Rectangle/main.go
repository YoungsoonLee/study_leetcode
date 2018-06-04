package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	fmt.Println(math.Sqrt(float64(area)))

	for i := int(math.Sqrt(float64(area))); i > 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}

	return []int{area, 1} // L>=w
}

func main() {
	a := 7
	fmt.Println(constructRectangle(a))
}
