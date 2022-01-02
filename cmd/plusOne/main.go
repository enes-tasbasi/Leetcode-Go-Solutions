package main

import "fmt"

func main() {
	digits := []int{9, 9, 9, 9}

	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	carry := -1

	for i := len(digits) - 1; i > -1; i-- {
		if carry == -1 || carry == 1 {
			if (digits[i]+1)%10 == 0 {
				carry = 1
				digits[i] = 0
			} else {
				digits[i]++
				carry = 0
			}
		}
	}

	if carry == 1 {
		newDigits := make([]int, len(digits)+1)
		for i, v := range digits {
			newDigits[i+1] = v
		}
		newDigits[0] = 1

		return newDigits
	}

	return digits
}

/*






 */
