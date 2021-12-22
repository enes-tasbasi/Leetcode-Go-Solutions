package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))
}

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	return sumArr(compare(nums))
}

func sumArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func max(nums map[int][]int) []int {
	max := -10001

	for k := range nums {
		if k > max {
			max = k
		}
	}

	return nums[max]
}

func compare(subArr []int) []int {
	if len(subArr) == 1 {
		return subArr
	}

	currSum := sumArr(subArr)

	leftSub := subArr[:len(subArr)/2]
	highestLeft := compare(leftSub)

	rightSub := subArr[len(subArr)/2:]
	highestRight := compare(rightSub)

	return max(map[int][]int{
		sumArr(highestLeft):  highestLeft,
		sumArr(highestRight): highestRight,
		currSum:              subArr,
	})
}
