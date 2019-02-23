package main

func combine(n int, k int) [][]int {

	temp := make([]int, k)

	for i := 1; i < n-k; i++ {
		temp = append(temp, i)
		subloop(i, k, &temp)
	}
}

func subloop(i, k int, t *[]int) {
	lv := (*)t[len(*t)-1]
	
}

func main() {
	n := 6
	k := 3
	combine(n, k)
}
