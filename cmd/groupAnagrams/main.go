package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func getCode(s string) int {
	t := 0
	for _, v := range s {
		t += int(v)
	}

	return t + len(s)
}

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
func groupAnagrams(strs []string) [][]string {
	st := [][]string{}
	hMap := map[string][]string{}

	for _, str := range strs {

		splitStr := strings.Split(str, "")
		sort.Strings(splitStr)
		complete := strings.Join(splitStr, "")

		if _, ok := hMap[complete]; ok {
			hMap[complete] = append(hMap[complete], str)
		} else {
			hMap[complete] = []string{str}
		}
	}

	for _, v := range hMap {
		local := []string{}
		local = append(local, v...)

		st = append(st, local)
	}

	return st
}
