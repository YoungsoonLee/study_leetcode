package main

/*
func findContentChildren(g []int, s []int) int {
	count := 0
	for si := 0; si < len(s); si++ {
		for gi := 0; gi < len(g); gi++ {
			if s[si] == g[gi] {
				count++
			}
		}

	}

	fmt.Println(count)
	return count
}
*/

/*
// g: greed factor, s: size cookie?
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	fmt.Println(g, s)

	var i, j, res int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	fmt.Println(res)
	return res
}
*/

func findContentChildren(g []int, s []int) int {
	totalCookies := 0
	content := 0
	for _, c := range s {
		totalCookies += c
	}

	for _, c := range g {
		if totalCookies > c {
			totalCookies -= c
			content++
		} else {
			break
		}
	}

	return content
}

func main() {
	a := []int{1, 2}
	s := []int{1, 2, 3}
	findContentChildren(a, s)
}
