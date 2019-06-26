package main

import (
	"fmt"
)

type MinStack struct {
	stack []item
}

// !!!!
type item struct {
	min, x int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	min := x
	if len(s.stack) > 0 && s.GetMin() < x {
		min = s.GetMin()
	}

	s.stack = append(s.stack, item{min: min, x: x})
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].x
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}

func main() {
	//ms := &Stack{val: []int{}, count: 0}
	ms := Constructor()
	ms.Push(3)
	//ms.Push(1)
	ms.Push(2)
	ms.Push(4)
	ms.Push(5)
	ms.Pop()
	//ms.GetMin()
	fmt.Println(ms.GetMin())
}
