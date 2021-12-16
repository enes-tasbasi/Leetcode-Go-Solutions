package main

import "fmt"

func main() {
	nums := []int{2, 2, 2, 3, 3, 4, 4}

	fmt.Println(nums)

	length := removeElement(nums, 2)

	for i := 0; i < length; i++ {
		fmt.Printf("%v ", nums[i])
	}

	fmt.Printf("\n")
	fmt.Printf("Lenght: %v\n", length)

	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			i--
			nums = nums[:len(nums)-1]
		}
	}

	return len(nums)
}
