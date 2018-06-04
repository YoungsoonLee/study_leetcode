package main

func generate(numRows int) [][]int {
	res := [][]int{}

	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}

	//fmt.Println(res)
	return res
}

func genNext(p []int) []int {
	r := make([]int, 1, len(p)+1) // 앞에다가 0은 널고, capacity 확보
	//fmt.Println(r)
	r = append(r, p...)

	//fmt.Println(r)

	for i := 0; i < len(r)-1; i++ {
		r[i] = r[i] + r[i+1]
		//fmt.Println(r)
	}

	return r
}

func main() {
	generate(5)
}
