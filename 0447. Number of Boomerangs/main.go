package main

func numberOfBoomerangs(points [][]int) int {

}

/*
func numberOfBoomerangs(points [][]int) int {
	res := 0

	size := len(points)
	fmt.Println("size: ", size)
	if size < 3 {
		return 0
	}

	for i := 0; i < size; i++ {
		// dMap [9]6은 점 [i]에서 거리가 3 인 점에 6 점이 있음을 나타냅니다.
		dMap := make(map[int]int, size)

		for j := 0; j < size; j++ {
			if i == j {
				continue
			}

			d := dSqure(points[i], points[j])
			if _, ok := dMap[d]; ok {
				dMap[d]++
			} else {
				dMap[d] = 1
			}
		}

		for _, v := range dMap {
			// 같은 거리에있는 점 v
			// 총 v * (v - 1) 정렬
			res += v * (v - 1)
		}
	}

	return res
}

func dSqure(p1, p2 []int) int {
	x := p2[0] - p1[0]
	y := p2[1] - p1[1]
	return x*x + y*y
}
*/

func main() {

}
