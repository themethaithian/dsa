package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	result := runningSum(nums)

	fmt.Printf("result: %+v\n", result)
}

func runningSum(nums []int) []int {
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])

	for i := 1; i < len(nums); i++ {
		result = append(result, result[i-1]+nums[i])
	}

	return result
}
