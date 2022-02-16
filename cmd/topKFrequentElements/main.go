package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(topKFrequent(data, k))
}

type KeyVal struct {
	num     int
	repeats int
}

type Arr []KeyVal

func (a Arr) Len() int           { return len(a) }
func (a Arr) Less(i, j int) bool { return a[i].repeats > a[j].repeats }
func (a Arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	arr := []KeyVal{}
	for k, v := range hMap {
		arr = append(arr, KeyVal{k, v})
	}

	sort.Sort(Arr(arr))

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, arr[i].num)
	}

	return res
}
