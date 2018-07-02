package main

import (
	"fmt"
	"log"
)

func main() {

	q := make([]int, 0)
	min := 1<<63 - 1
	re := make([]int, 0)

	var op string
	var num int

	for {
		fmt.Scanf("%s %d", &op, &num)
		if op == "add" {
			add(num, &q, &min)
		}

		if op == "min" {
			re = append(re, min)
		}

		if op == "remove" {
			re = append(re, remove(&q, &min))
		}

		if op == "exit" {
			break
		}
	}

	for _, r := range re {
		fmt.Println(r)
	}
}

func add(n int, q *[]int, m *int) {
	*q = append(*q, n)

	if len(*q) == 1 {
		*m = n
	} else {
		for i := 0; i < len((*q))-1; i++ {
			//fmt.Println(getMin((*q)[i], (*q)[i+1]), (*q)[i], (*q)[i+1])
			*m = min((*q)[i], (*q)[i+1])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func remove(q *[]int, m *int) int {
	if len(*q) == 0 {
		log.Fatalln("Exception not exists")
	}

	x := (*q)[0]
	// Discard top element
	*q = (*q)[1:]

	if len(*q) == 1 {
		*m = (*q)[0]
	} else {
		for i := 0; i < len((*q))-1; i++ {
			//fmt.Println(getMin((*q)[i], (*q)[i+1]), (*q)[i], (*q)[i+1])
			*m = min((*q)[i], (*q)[i+1])

		}
	}

	return x
}
