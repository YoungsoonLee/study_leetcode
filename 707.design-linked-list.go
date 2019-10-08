/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 *
 * https://leetcode.com/problems/design-linked-list/description/
 *
 * algorithms
 * Easy (20.92%)
 * Likes:    368
 * Dislikes: 531
 * Total Accepted:    35.4K
 * Total Submissions: 169.7K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * Design your implementation of the linked list. You can choose to use the
 * singly linked list or the doubly linked list. A node in a singly linked list
 * should have two attributes: val and next. val is the value of the current
 * node, and next is a pointer/reference to the next node. If you want to use
 * the doubly linked list, you will need one more attribute prev to indicate
 * the previous node in the linked list. Assume all nodes in the linked list
 * are 0-indexed.
 * 
 * Implement these functions in your linked list class:
 * 
 * 
 * get(index) : Get the value of the index-th node in the linked list. If the
 * index is invalid, return -1.
 * addAtHead(val) : Add a node of value val before the first element of the
 * linked list. After the insertion, the new node will be the first node of the
 * linked list.
 * addAtTail(val) : Append a node of value val to the last element of the
 * linked list.
 * addAtIndex(index, val) : Add a node of value val before the index-th node in
 * the linked list. If index equals to the length of linked list, the node will
 * be appended to the end of linked list. If index is greater than the length,
 * the node will not be inserted. If index is negative, the node will be
 * inserted at the head of the list.
 * deleteAtIndex(index) : Delete the index-th node in the linked list, if the
 * index is valid.
 * 
 * 
 * Example:
 * 
 * 
 * MyLinkedList linkedList = new MyLinkedList();
 * linkedList.addAtHead(1);
 * linkedList.addAtTail(3);
 * linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
 * linkedList.get(1);            // returns 2
 * linkedList.deleteAtIndex(1);  // now the linked list is 1->3
 * linkedList.get(1);            // returns 3
 * 
 * 
 * Note:
 * 
 * 
 * All values will be in the range of [1, 1000].
 * The number of operations will be in the range of [1, 1000].
 * Please do not use the built-in LinkedList library.
 * 
 * 
 */

// @lc code=start
type MyLinkedList struct {
	size int
	head, tail *node
}

type node {
	val int
	next *node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	t := &node{}
	h := &node{next: t}
	return MyLinkedList{
		head: h,
		tail: t,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.size <= index {
		return -1
	}

	i, cur := 0, this.head.next
	for i < index {
		i++
		cur = cur.next
	}

	return cur.val

}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	nd := &node{val: val}
	nd.next = this.head.next
	this.head.next = nd
	this,size++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.tail.val = val // !!

	nd := &node{}
	this.tail.next = nd
	this.tail = nd
	this.size++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	switch {
		case index < 0 || this.size < index:
			return
		case index == 0:
			this.AddAtHead(val)
			return
		case index == this.size:
			this.AddAtTail(val)
			return
	}

	i, cur  := -1, this.head
	for i+1 < index{
		i++
		cur = cur.next
	}

	nd := &node{val: val}
	nd.next = cu.next
	cur.next = nd
	this.size++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || l.size <= index {
		return
	}

	i, cur := -1, this.head
	for i+1 < index {
		i++
		cur = cur.next
	}

	cur.next = cur.next.next
	this.size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

