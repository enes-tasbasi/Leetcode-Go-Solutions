package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 2, 7, 8, 8}

	newLength := removeDuplicates(nums)

	for i := 0; i < newLength; i++ {
		fmt.Printf("%v ", nums[i])
	}
	fmt.Printf("\n")

	fmt.Printf("Length: %v\nElements:%v\n", nums, newLength)
}

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	fmt.Println(nums)

	return len(nums)
}
