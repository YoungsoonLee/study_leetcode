package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {
	up := strings.Count(moves, "U")
	down := strings.Count(moves, "D")
	left := strings.Count(moves, "L")
	right := strings.Count(moves, "R")

	return up == down && left == right
}

func judgeCircle_my(moves string) bool {
	move := make(map[string]int, 4)
	move["R"] = 1
	move["L"] = -1
	move["U"] = 1
	move["D"] = -1

	fmt.Println(move)

	cm := 0

	for _, c := range moves {
		cm += move[string(c)]
	}

	if cm == 0 {
		return true
	}

	return false

}

func main() {
	a := "UD"
	fmt.Println(judgeCircle(a))
}
