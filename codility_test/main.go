package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	l := len(A)
	sort.Ints(A)
	fmt.Println(A)

	t := 0
	h := l - 1
	minAbsSum := 0
	for t <= h {
		cSum := A[t] + A[h]
		minAbsSum = max(minAbsSum, abs(cSum))

		if cSum <= 0 {
			t++
		} else {
			h--
		}
	}

	fmt.Println(minAbsSum)
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Solution_2(A []int) int {
	//l := len(A)
	result := 0
	pair := make([][]int, 0)
	for i, a := range A {
		pair = append(pair, []int{a, i})
	}

	fmt.Println(pair)

	sort.Slice(pair[:], func(i, j int) bool {
		for x := range pair[i] {
			if pair[i][x] == pair[j][x] {
				continue
			}
			return pair[i][x] < pair[j][x]
		}
		return false
	})

	fmt.Println(pair)

	for i := 1; i < len(pair)-1; i++ {
		fmt.Println(pair[i-1][1], pair[i][1])
		result = max(result, abs(abs(pair[i][1])-abs(pair[i-1][1])))
	}

	fmt.Println(result)
	return 0

}

func Solution_1(A []int) int {
	// write your code in Go 1.4
	if len(A) == 0 || len(A) == 1 {
		return -1
	}

	sort.Ints(A)

	fmt.Println(A)

	pre := A[0]
	curCount := 1
	preCount := 0
	result := 0

	for i := 1; i < len(A); i++ {

		fmt.Println(pre, curCount, preCount, result)

		if A[i] == pre {
			curCount++
			result += preCount + curCount - 1
		} else {
			result += curCount
			preCount = curCount
			curCount = 1
		}

		pre = A[i]
	}

	fmt.Println(result)
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		log.Printf("Failed to read: %v", scanner.Err())
		return
	}
	line := scanner.Bytes()
	s := strings.Split(string(line), " ")
	fmt.Println("bb: ", string(line))
	fmt.Println("s: ", s[0])

}
