package main

func numTrees(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	res := 0
	// 대칭이므로 절반 만
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-1)
	}

	res *= 2 // ????
	// n이 홀수 인 경우를 처리합니다.
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}

	return res
}

func main() {

}
