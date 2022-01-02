package main

import "fmt"

func main() {
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	arr := []string{}
	str := -1

	if len(s) == 1 {
		if s == " " {
			return 0
		} else {
			return 1
		}
	}

	for i, v := range s {
		if v == ' ' {
			if str > -1 {
				arr = append(arr, s[str:i])
			}
			str = -1
		} else if i == len(s)-1 {
			if str > -1 {
				arr = append(arr, s[str:])
			}
		} else {
			if str < 0 {
				str = i
			}
		}
	}

	max := 0
	for _, v := range arr {
		if len(v) > max {
			max = len(v)
		}
	}

	return max
}
