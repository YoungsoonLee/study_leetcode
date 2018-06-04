package main

type MyStack struct {
	a, b *Queue
}

type Queue struct {
	nums []int
}

func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

// push onto stack
func (ms *MyStack) Push(x int) {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a // ?? why?? - 임시 보관으로 b에 값이 있을 수 있다.
	}
	ms.a.enqueue(x)
}

// pop stack !!!!
func (ms *MyStack) Pop() int {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a // ?? why??
	}

	for ms.a.Len() > 1 { // 하나 남을때 까지...
		ms.b.enqueue(ms.a.dequeue())
	}

	return ms.a.dequeue()
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Top
func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res) // remake stack
	return res
}

func (ms *MyStack) Empty() bool {
	return (ms.a.Len() + ms.b.Len()) == 0
}

// enqueue queue
func (q *Queue) enqueue(n int) {
	q.nums = append(q.nums, n)
}

// dequeue queue
func (q *Queue) dequeue() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) isEmpty() bool {
	return len(q.nums) == 0
}

func main() {
	/**
	 * Your MyStack object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Push(x);
	 * param_2 := obj.Pop();
	 * param_3 := obj.Top();
	 * param_4 := obj.Empty();
	 */
}
