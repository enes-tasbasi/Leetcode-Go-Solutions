package main

import "fmt"

func main() {

	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for a := i + 1; a < len(nums); a++ {
			if nums[i]+nums[a] == target {
				return []int{i, a}
			}
		}
	}

	return []int{}
}
