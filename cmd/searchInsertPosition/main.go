package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}

	fmt.Println(nums)

	fmt.Println(searchInsert(nums, 5))
}

/*
 * Given a sorted array of distinct integers and a target value,
 * return the index if the target is found.
 * If not, return the index where it would be if it were inserted in order.
 */
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
