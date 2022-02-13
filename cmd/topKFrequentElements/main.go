package main

import (
	"fmt"

	heap "github.com/etasbasi/leetcode/m/v2/utils"
)

func main() {
	data := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(topKFrequent(data, k))
}

// TODO: fails because it is not fast enough on large data, find a way to use a priority queue.
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
func topKFrequent(nums []int, k int) []int {
	hMap := map[int]int{}

	for _, v := range nums {
		if _, ok := hMap[v]; ok {
			hMap[v]++
		} else {
			hMap[v] = 1
		}
	}

	pq := heap.MaxHeap{}

	for _, v := range hMap {
		pq.Push(v)
	}

	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res, pq.Pop())
	}

	return res
}
