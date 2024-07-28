package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	result := longestOnes(nums, k)

	fmt.Printf("result: %d\n", result)

	nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 3
	result = longestOnes(nums, k)

	fmt.Printf("result: %d\n", result)
}

func longestOnes(nums []int, k int) int {
	left := 0
	ans := 0
	curr := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > k {
			if nums[left] == 0 {
				curr--
			}
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
