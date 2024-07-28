package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)

	fmt.Printf("%+v", result)
}

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	cur := len(result) - 1

	i := 0
	j := len(nums) - i - 1

	for {
		if nums[i]*nums[i] >= nums[j]*nums[j] {
			result[cur] = nums[i] * nums[i]
			i++
		} else {
			result[cur] = nums[j] * nums[j]
			j--
		}

		if cur == 0 {
			break
		}

		cur--
	}

	return result
}
