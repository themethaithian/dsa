package main

import "fmt"

func main() {
	s := "110110111"
	result := findLength(s)

	fmt.Printf("result: %d", result)
}

func findLength(s string) int {
	var left, curr, ans int

	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			curr++
		}

		for curr > 1 {
			if s[left] == '0' {
				curr--
			}

			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
