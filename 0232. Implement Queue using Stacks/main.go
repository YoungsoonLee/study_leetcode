package main

import "fmt"

type MyQueue struct {
	s, tmp *Stack
}

func Constructor() MyQueue {
	return MyQueue{s: NewStack()}
}

// enqueue
func (q *MyQueue) enqueue(x int) {
	/*
		if q.s.Len() == 0 {
			q.s, q.tmp = q.tmp, q.s
		}
	*/
	q.s.Push(x)
}

// dequeue
func (q *MyQueue) dequeue() int {
	/*
		if q.s.Len() == 0 {
			q.s, q.tmp = q.tmp, q.s
		}
		for q.s.Len() > 1 {
			q.tmp.Push(q.s.Pop())
		}

		return q.s.Pop()
	*/
	if q.tmp.Len() == 0 {
		for q.s.Len() > 0 {
			q.tmp.Push(q.s.Pop())
		}
	}

	return q.tmp.Pop()
}

// peek - finr front element
func (q *MyQueue) peek() int {
	/*
		res := q.dequeue()
		q.s.Push(res)
		q.s = q.tmp
		return res
	*/
	res := q.dequeue()
	q.tmp.Push(res)
	return res
}

// empty
func (q *MyQueue) Empty() bool {
	return (q.s.Len() + q.tmp.Len()) == 0
}

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// push stack
func (s *Stack) Push(x int) {
	s.nums = append(s.nums, x)
}

// pop stack
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
