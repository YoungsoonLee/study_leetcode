package main

import (
	"fmt"
	"sort"
)

func minimumTotal(triangle [][]int) int {
	/** ???? **/
	/* below is not a answer */
	sum := 0
	for i := 0; i < len(triangle); i++ {
		sort.Ints(triangle[i])
		sum += triangle[i][0]
	}

	fmt.Println("sum: ", sum)
	return sum
}

func main() {

}
