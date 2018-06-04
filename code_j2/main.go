package main

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func Solution(T *Tree) int {
	// write your code in Go 1.4
	preValue := T.Val
	resultCount := 1
	search(T, &preValue, &resultCount)
	return resultCount
}

func search(root *Tree, pv *int, rc *int) {
	if root == nil {
		return
	}

	search(root.Left, pv, rc)

	if comp(*pv, root.Val) {
		*rc++
	}

	*pv = root.Val

	search(root.Right, pv, rc)
}

func comp(prev, curv int) bool {
	if prev < curv {
		return true
	}
	return false
}

func main() {

	t2l := &Tree{Val: 20, Left: nil, Right: nil}
	t2r := &Tree{Val: 21, Left: nil, Right: nil}

	t3l := &Tree{Val: 1, Left: nil, Right: nil}

	t1l := &Tree{Val: 3, Left: t2l, Right: t2r}
	t1r := &Tree{Val: 10, Left: t3l, Right: nil}

	t1Root := &Tree{Val: 5, Left: t1l, Right: t1r}

	Solution(t1Root)

}
