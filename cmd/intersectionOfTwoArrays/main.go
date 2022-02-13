package main

import "fmt"

func main() {
	nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
}

// Given two integer arrays nums1 and nums2,
// return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.
func intersection(nums1 []int, nums2 []int) []int {
	hMap := map[int]bool{}
	res := map[int]bool{}

	for i := 0; i < len(nums1); i++ {
		hMap[nums1[i]] = true
	}

	for a := 0; a < len(nums2); a++ {
		if _, ok := hMap[nums2[a]]; ok {
			res[nums2[a]] = true
		}
	}

	resArr := []int{}
	for i := range res {
		resArr = append(resArr, i)
	}
	return resArr
}
