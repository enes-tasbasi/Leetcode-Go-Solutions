package main

import "fmt"

func main() {

	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		remainder := target - nums[i]

		if _, ok := nMap[remainder]; ok {
			return []int{i, nMap[remainder]}
		}

		nMap[nums[i]] = i
	}

	return []int{}
}
