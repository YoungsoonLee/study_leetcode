/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (40.34%)
 * Likes:    3071
 * Dislikes: 298
 * Total Accepted:    517.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 * 
 * Output
 * [null,null,null,null,-3,null,0,-2]
 * 
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 * 
 * 
 */

// @lc code=start
type MinStack struct {
    stack []item
}

type item struct {
    min, x int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    min := x
    if len(this.stack) >0 && this.GetMin() < x {
        min = this.GetMin()
    }
    this.stack = append(this.stack, item{min: min, x: x})
    //this.Print()
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
    //this.Print()
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1].x
}


func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack)-1].min
}

func (this *MinStack) Print() {
    fmt.Println(this.stack)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

