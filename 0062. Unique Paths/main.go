package main

import "fmt"

func uniquePaths(m int, n int) int {
	var routeToEnd = make([][]int, n)
	fmt.Println(routeToEnd)

	for i := range routeToEnd {
		routeToEnd[i] = make([]int, m)
	}
	fmt.Println(routeToEnd)

	for i := 0; i < n; i++ {
		routeToEnd[i][0] = 1
	}
	fmt.Println(routeToEnd)

	for i := 0; i < m; i++ {
		routeToEnd[0][i] = 1
	}
	fmt.Println(routeToEnd)

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			routeToEnd[i][j] = routeToEnd[i-1][j] + routeToEnd[i][j-1]
		}
	}
	fmt.Println(routeToEnd)
	fmt.Println(routeToEnd[n-1][m-1])
	return routeToEnd[n-1][m-1]
}

func main() {
	m, n := 3, 2
	uniquePaths(m, n)

}
