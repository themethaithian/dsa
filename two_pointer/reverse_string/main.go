package main

import "fmt"

func main() {
	s := []byte{'H'}
	reverseString(s)
	fmt.Printf("%+v", s)
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
