package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Hello, world!\n")
	strings := []string{"Postfix", "Postman", "Pomodora"}

	fmt.Printf("longcommonprefix: %s\n", longestCommonPrefix(strings))
}

func longestCommonPrefix(strs []string) string {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	prefix := ""

	shortestWordLength := MaxInt
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < shortestWordLength {
			shortestWordLength = len(strs[i])
		}
	}

	fmt.Printf("shortest: %d\n", shortestWordLength)

	for i := 0; i < shortestWordLength; i++ {

		currChar := string(strs[0][i])
		for a := 0; a < len(strs); a++ {
			if strings.Compare(string(strs[a][i]), currChar) != 0 {
				return prefix
			}
		}
		prefix += currChar
	}
	return prefix
}
