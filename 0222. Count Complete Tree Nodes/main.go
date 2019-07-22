package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		count++

		dfs(root.Left)
		dfs(root.Right)

	}

	dfs(root)

	return count
}

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	s := []int{1, 2, 3}
	b := s

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println("--1----------")

	b[0] = 999
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println("--2----------")

	var c []int
	c = make([]int, 3)
	copy(c, s) // 새롭게 만들어서 복사니깐...
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("---3---------")

	c[0] = 888
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("---4---------")

	a := Student{"aaa", 20, 10}
	d := a

	d.age = 30

	fmt.Println(a)
	fmt.Println(d)

	fmt.Println("---5---------")

	y := "a"
	z := y
	z = "b"
	fmt.Println(y)
	fmt.Println(z)
}
