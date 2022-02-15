package main

import (
	"fmt"
	"strings"
)

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func hasAll(a, b string) bool {
	remainA := a
	if len(a) != len(b) {
		return false
	}

	for _, letter := range b {
		index := strings.IndexRune(remainA, letter)
		if index == -1 {
			return false
		}

		remainA = strings.Replace(remainA, string(letter), "", 1)
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	st := [][]string{}

	for i := 0; i < len(strs); i++ {

		hasWord := false

		for a := 0; a < len(st); a++ {
			if len(st[a]) == 0 {
				continue
			}

			if hasAll(st[a][0], strs[i]) {
				hasWord = true

				st[a] = append(st[a], strs[i])
				break
			}
		}

		if !hasWord {
			st = append(st, []string{strs[i]})
		}
	}

	return st
}
