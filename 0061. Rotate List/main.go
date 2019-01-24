package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//var root = new(ListNode)

func rotateRight(head *ListNode, k int) *ListNode {
	//fmt.Println("k: ", k)

	if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			// 处理 k 大于 list 的长度的情况
			// i+1 就是 list 的长度

			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}

	fmt.Println("tail: ", tail)

	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	//fmt.Println("tail: ", tail)
	//fmt.Println("newTail: ", newTail)

	traverse(tail)
	traverse(newTail)

	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head

	// traverse(newHead)

	return newHead

	/*
		prev := new(ListNode)
		temp := new(ListNode)

		for i := 0; i < k; i++ {
			for head != nil {
				prev.Val = head.Val

				if head.Next == nil {
					temp.Val = head.Val
					prev.Next = nil
				}
				head = head.Next
			}
		}

		temp.Next = root

		return root
	*/
}

func addNode(l *ListNode, v int) int {
	if l == nil {
		// root
		//temp := &ListNode{v, nil}
		//root = temp
		l = &ListNode{v, nil}
		return 0
	}

	if v == l.Val {
		// do nothing
		return -1
	}

	if l.Next == nil {
		l.Next = &ListNode{v, nil}
		return -2
	}

	return addNode(l.Next, v)
}

func traverse(l *ListNode) {
	if l == nil {
		fmt.Println("--> Empty list!")
		return
	}

	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}

	fmt.Println()
}

func main() {
	root := &ListNode{1, nil}
	//addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	addNode(root, 5)
	traverse(root)
	root = rotateRight(root, 2)
	traverse(root)
}
