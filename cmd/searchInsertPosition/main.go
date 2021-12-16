package main

import "fmt"

func main() {
	nums := []int{1, 3, 5}

	fmt.Println(nums)

	fmt.Println(searchInsert(nums, 4))
}

/*
 * Given a sorted array of distinct integers and a target value,
 * return the index if the target is found.
 * If not, return the index where it would be if it were inserted in order.
 */
func searchInsert(nums []int, target int) int {
	lowerBound := 0
	upperBound := len(nums)

	for {
		if upperBound-lowerBound == 1 {
			if nums[lowerBound] == target {
				return lowerBound
			} else {
				if target > nums[lowerBound] {
					return lowerBound + 1
				} else {
					if lowerBound-1 < 0 {
						return 0
					}

					return lowerBound - 1
				}
			}
		}

		newDivident := lowerBound + (upperBound / 2)
		if newDivident >= len(nums) {
			return len(nums)
		}
		if target == nums[newDivident] {
			return lowerBound + (upperBound / 2)
		}
		if target < nums[newDivident] {
			upperBound = lowerBound + (upperBound / 2)
		} else {
			lowerBound = lowerBound + (upperBound / 2)
		}
	}
}

/*
nums: [1, 2, 3] target: 3

arrSlice <- nums [1, 2, 3] | [2, 3]
while true:
	if nums.length == 1 { // false | false
		if nums[0] == target
	}
	if target == nums[nums.length / 2] { // false | true
		return nums.length / 2 // 1
	}
	if	target < nums[nums.length / 2] { // false
		arrSlice <- nums.slice(0, nums.length / 2)
	} else {
		arrSlice <- nums.slice(nums.length / 2, nums.length) // [2, 3]
	}


nums: [1, 2, 3]
target: 1

lowerBound: 0
upperBound: nums.length // 3 | 1

while true:

	if upperBound - lowerBound == 1 { // 0, false | 1, true
		if nums[lowerBound] == target {
			return lowerBound
		} else {
			if target > nums[lowerBound] {
				return lowerBound + 1
			} else {
				return lowerBound - 1
			}
		}
	}

	if target == nums[lowerBound + (upperBound / 2) ] { // 2, false
		return lowerBound + (upperBound / 2)
	}
	if	target < nums[lowerBound + (upperBound / 2)] { // true
		upperBound: upperBound / 2 // upperBound: 1
	} else {
		lowerBound: upperBound / 2
	}


*/
