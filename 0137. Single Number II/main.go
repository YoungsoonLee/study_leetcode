package main

import "sort"

func singleNumber2(nums []int) int {
	// ref : https://cloud.tencent.com/developer/article/1131945

	// 두 번 나타나면 조금씩 사용하십시오.
	// 예를 들어, 1, 초기 0
	// 5가 먼저 나타나면 ones = 5
	// 5가 두 번째로 나타나면 1이 0으로 지워지며 다른 숫자를 처리 할 수 ​​있음을 나타냅니다.
	// 그래서, 마침내, ones! = 0 일 경우, 한 번만 나타나는 숫자를 기록합니다.

	// 수식은 ones = ones xor i입니다.

	// 그러면 세 번이면 Shannon의 정리를 2 비트로 기록해야합니다.

	// 5가 처음 나타날 때 ones = 5, twos = 0,이 숫자가 기록됨
	// 5가 두 번째로 나타날 때 ones = 0, twos = 5, twos가이 수를 기록합니다.
	// 5가 세 번째 시간 인 경우 ones = 0, twos = 0, 모두 지워지면 다른 숫자를 처리 할 수 ​​있습니다.
	// 그래서, 한 번 나타나는 숫자가 있다면, 1이 있고, 2 번 있고, 두 개가있다.

	// 공식, 위의 두 번, 비어있는 수식은 ones = ones xor i입니다.
	// 3 번째에있는 것들은 0과 같고 그 다음에 두 개가 참이된다. 그래서 두 개가 될 수있다 : ones = (ones xor i) & ~ twos
	// 따라서 전체 수식은 다음과 같습니다.
	// ones = (ones xor i) & ~ twos
	// twos = (twos xor i) & ~ ones
	ones, twos := 0, 0
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & ^twos
		twos = (twos ^ nums[i]) & ^ones
	}
	return ones
}

func singleNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 {
		return nums[0]
	}

	i := 0
	for i < len(nums)-3 {

		if nums[i] != nums[i+1] {
			return nums[i]
		}

		i = i + 3
	}

	if i == len(nums)-1 {
		return nums[i]
	}

	return 0
}

func main() {

}
