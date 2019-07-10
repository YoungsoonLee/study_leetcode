package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{stack: stack}
	res.push(root)
	return res
}

func (it *BSTIterator) push(root *TreeNode) {
	for root != nil {
		it.stack = append(it.stack, root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	size := len(this.stack)
	var top *TreeNode
	this.stack, top = this.stack[:size-1], this.stack[size-1]
	this.push(top.Right)
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {

}
