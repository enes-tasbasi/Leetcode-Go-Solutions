package main

import "fmt"

func main() {
	s := " "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	left := 0
	right := 0
	lastMax := 0

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			left = i

			for right = i + 1; right < len(s); right++ {
				if s[right] == ' ' {
					break
				}
			}
			i = right

		}
	}

	lastMax = right - left

	return lastMax
}
