package main

import "fmt"

func main() {

	q := make([]int, 0)
	min := 1<<63 - 1

	var op string
	var num int

	for {
		fmt.Scanf("%s %d", &op, &num)
		if op == "add" {
			add(num, &q, &min)
		}

		if op == "min" {
			fmt.Println(min)
		}

		if op == "remove" {
			remove(&q, &min)
		}

		if op == "exit" {
			break
		}
	}

	//fmt.Println(min)
	//add(5, q, &min)
}

func add(n int, q *[]int, m *int) {
	*q = append(*q, n)

	if len(*q) == 1 {
		*m = n
	} else {
		for i := 0; i < len((*q))-1; i++ {
			//fmt.Println(getMin((*q)[i], (*q)[i+1]), (*q)[i], (*q)[i+1])
			*m = getMin((*q)[i], (*q)[i+1])
		}
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func remove(q *[]int, m *int) {
	x := (*q)[0]
	// Discard top element
	*q = (*q)[1:]

	if len(*q) == 1 {
		*m = (*q)[0]
	} else {
		for i := 0; i < len((*q))-1; i++ {
			//fmt.Println(getMin((*q)[i], (*q)[i+1]), (*q)[i], (*q)[i+1])
			*m = getMin((*q)[i], (*q)[i+1])

		}
	}

	fmt.Println(x)
}
