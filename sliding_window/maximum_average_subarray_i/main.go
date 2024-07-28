package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(nums, k)

	fmt.Printf("result: %v\n", result)
}

func findMaxAverage(nums []int, k int) float64 {
	curr := 0
	for i := 0; i < k; i++ {
		curr += nums[i]
	}

	ans := curr 

	for right := k; right < len(nums); right++ {
		curr += nums[right] - nums[right-k]

		ans = max(ans, curr)
	}

	return float64(ans) / float64(k)
}
