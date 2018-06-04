package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	// inpu check 32bit signed integer overflow
	if x > math.MaxInt32 || x < -math.MaxInt32 {
		return 0
	}

	s := strconv.Itoa(int(math.Abs(float64(x))))
	rs := make([]string, len(s), len(s))

	for i := len(s) - 1; i >= 0; i-- {
		rs = append(rs, string(s[i]))
	}

	ri, _ := strconv.Atoi(strings.Join(rs, ""))

	if x < 0 {
		if -1*ri > math.MaxInt32 || -1*ri < -math.MaxInt32 {
			return 0
		} else {
			return -1 * ri
		}

	} else {
		if ri > math.MaxInt32 || ri < -math.MaxInt32 {
			return 0
		} else {
			return ri
		}
	}
}

func reverse2(x int) int {
	// inpu check 32bit signed integer overflow
	if x > math.MaxInt32 || x < -math.MaxInt32 {
		return 0
	}

	rn := 0
	for x != 0 {
		rn = rn*10 + x%10
		x /= 10
		fmt.Println(x, rn)
	}

	// output check 32bit signed integer overflow
	if rn > math.MaxInt32 || rn < -math.MaxInt32 {
		return 0
	}
	return rn

}

func main() {
	a := -123
	fmt.Println(reverse2(a))
}
