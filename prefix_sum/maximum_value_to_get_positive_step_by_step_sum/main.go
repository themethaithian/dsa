package main

import (
	"fmt"
)

func main() {
	nums := []int{-3, 2, -3, 4, 2}
	result := minStartValue(nums)

	fmt.Printf("result: %d\n", result)
}

func minStartValue(nums []int) int {
	curSum := nums[0]
	minSum := nums[0]

	for i := 1; i < len(nums); i++ {
		curSum += nums[i]
		minSum = min(minSum, curSum)
	}

	if minSum >= 0 {
		return 1
	}

	return -minSum + 1
}
