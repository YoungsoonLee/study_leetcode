package main

import "fmt"

type MyQueue struct {
	s, tmp *Stack
}

func Constructor() MyQueue {
	return MyQueue{s: NewStack(), tmp: NewStack()}
}

func (q *MyQueue) Push(x int) {
	q.s.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.tmp.Len() == 0 {
		for q.s.Len() > 0 { // !!!! loop
			q.tmp.Push(q.s.Pop())
		}
	}
	return q.tmp.Pop()
}

func (q *MyQueue) Peek() int {
	res := q.Pop()
	q.tmp.Push(res)
	return res
}

func (q *MyQueue) Empty() bool {
	return (q.s.Len() + q.tmp.Len()) == 0
}

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (s *Stack) Push(x int) {
	s.nums = append(s.nums, x)
}

func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

// top stack - seek top
func (s *Stack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack) Len() int {
	return len(s.nums)
}

func main() {
	fmt.Println("")
}
