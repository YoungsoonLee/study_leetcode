package main

func partition(s string) [][]string {
	result := [][]string{}
	current := make([]string, 0, len(s))

	dfs(s, 0, current, &result)
	return result
}

func dfs(s string, i int, cur []string, result *[][]string) {
	if i == len(s) {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for j := i; j < len(s); j++ {
		// i == 0 일 때,
		// len (cur [0])에 따라 res를 나눕니다.
		// 등등
		if part(s[i : j+1]) {
			dfs(s, j+1, append(cur, s[i:j+1]), result) // !!!
		}
	}
}

// s는 회문이며 true를 반환합니다.
func part(s string) bool {
	if len(s) <= 1 {
		return true
	}

	a, b := 0, len(s)-1
	for a < b {
		if s[a] != s[b] {
			return false
		}
		a++
		b--
	}
	return true
}

/*
func partition(s string) [][]string {
	res := make([][]string, 0)
	temp := make([]string, 0)

	for c := 0; c < len(s); c++ {
		//temp := []string{}

		for n := c + 1; n < len(s); n++ {
			if s[c] == s[n] {
				s := string(s[c]) + string(s[n])
				//temp = append(temp, s)
				temp = append(temp, s)
			} else {
				break
			}
		}
		fmt.Println(c, len(s))
		if c+1 == len(s) {
			//temp = append(temp, string(s[c]))
			temp = append(temp, string(s[c]))
			res = append(res, temp)
			temp = []string{}
		}
	}

	// final
	st := []string{}
	for i := 0; i < len(s); i++ {
		st = append(st, string(s[i]))
	}
	res = append(res, st)

	fmt.Println(res)
	return res
}
*/

func main() {
	s := "aab"
	partition(s)
}
