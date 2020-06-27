/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 *
 * https://leetcode.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (42.22%)
 * Likes:    630
 * Dislikes: 570
 * Total Accepted:    177.4K
 * Total Submissions: 398.8K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * Implement the following operations of a stack using queues.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * empty() -- Return whether the stack is empty.
 * 
 * 
 * Example:
 * 
 * 
 * MyStack stack = new MyStack();
 * 
 * stack.push(1);
 * stack.push(2);  
 * stack.top();   // returns 2
 * stack.pop();   // returns 2
 * stack.empty(); // returns false
 * 
 * Notes:
 * 
 * 
 * You must use only standard operations of a queue -- which means only push to
 * back, peek/pop from front, size, and is empty operations are valid.
 * Depending on your language, queue may not be supported natively. You may
 * simulate a queue by using a list or deque (double-ended queue), as long as
 * you use only standard operations of a queue.
 * You may assume that all operations are valid (for example, no pop or top
 * operations will be called on an empty stack).
 * 
 * 
 */

// @lc code=start
type MyStack struct {
	a, b *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{a: NewQueue(), b: NewQueue()}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    if this.a.len() == 0 {
		this.a, this.b = this.b this.a
	}
	this.a.Push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Len() == 0 {
		this.a, this.b == this.b, this.a
	}

	for this.a.Len() > 1 {
		this.b.Push(this.a.Pop())
	}

	return this.a.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	res := this.a.Pop()
	this.Push(res)
	return res
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return (this.a.Len()+this.b.Len()) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

 type Queue struct {
	 nums []int
 }

 func NewQueue() *Queue {
	return &Queue{nums: []int{}}
 }

 func (q *Queue) Push(n int) {
	 q.nums = append(q.nums, n)
 }

 func (q *Queue) Pop() int {
	 t := q.nums[0]
	 q.nums = q.num[1:]
	 return t
 }

 func (q *Queue) Len() int {
	 return len(q.nums)
 }

 func (q *Queue) Empty() bool {
	 reutnn q.Len() == 0
 }

// @lc code=end

