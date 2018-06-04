package main

import "fmt"

func main() {
	pair := make([][]int, 0)
	getlocation(0, 1, 1, &pair)
	getlocation(1, 2, 2, &pair)
}

func getlocation(x, w, h int, p *[][]int) {
	rbXY := []int{x + w, 0}
	rtXY := []int{x + w, h}
	lbXY := []int{x, 0}
	ltXY := []int{x, h}

	*p = append(*p, rbXY)
	*p = append(*p, rtXY)
	*p = append(*p, lbXY)
	*p = append(*p, ltXY)

	fmt.Println(*p)
}

func unique(p *[][]int) {
	//keys := make(map[int]bool)

	for k, v := range *p {
		fmt.Println(k, v)
	}

}
