package main

import "fmt"

func main() {
	nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	result := getAverages(nums, k)

	fmt.Printf("result: %+v\n", result)
}

func getAverages(nums []int, k int) []int {
	result := make([]int, len(nums))
	prefixSumNums := make([]int, 0, len(nums))
	prefixSumNums = append(prefixSumNums, nums[0])

	for i := 1; i < len(nums); i++ {
		prefixSumNums = append(prefixSumNums, prefixSumNums[i-1]+nums[i])
	}

	for i := 0; i < len(nums); i++ {
		if i < k || i > len(nums)-k-1 {
			result[i] = -1
		} else {
			result[i] = (prefixSumNums[i+k] - prefixSumNums[i-k] + nums[i-k]) / (k*2 - 1)
		}
	}

	return result
}
