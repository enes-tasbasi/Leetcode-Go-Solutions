package main

import "fmt"

func main() {

	str := "Hello there"

	i := strStr(str, "res")

	fmt.Println(i)
}

// Return the index of the first occurrance of the needle inside the haystack
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	if needle == "" {
		return 0
	}

	var compare = func(first, second string) bool {
		if len(first) != len(second) {
			return false
		}
		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				return false
			}
		}

		return true
	}

	for i, v := range haystack {
		if v == rune(needle[0]) {
			if len(haystack[i:]) < len(needle) {
				return -1
			}
			if compare(haystack[i:i+len(needle)], needle) {
				return i
			}
		}
	}

	return -1
}
