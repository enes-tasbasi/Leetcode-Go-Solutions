package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))
}

func max(nums ...int) int {
	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return max
}

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}
