package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	val []int
	min int
}

func Constructor() Stack {
	return Stack{}
}

func (s *Stack) push(item int) {
	s.val = append(s.val, item)
}

func (s *Stack) pop() (int, error) {
	x := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1] // 개수!!!
	return x, nil
}

func (s *Stack) top() (int, error) {

	x := s.val[len(s.val)-1]
	return x, nil
}

func (s *Stack) GetMin() int {
	minv := make([]int, len(s.val))
	copy(minv, s.val)
	//fmt.Println("minv: ", minv)
	sort.Sort(sort.IntSlice(minv))

	return minv[0]
}

func main() {
	//ms := &Stack{val: []int{}, count: 0}
	ms := Constructor()
	ms.push(3)
	ms.push(1)
	ms.push(2)
	ms.push(4)
	ms.push(5)
	ms.pop()
	//ms.GetMin()
	fmt.Println(ms.GetMin())
}
