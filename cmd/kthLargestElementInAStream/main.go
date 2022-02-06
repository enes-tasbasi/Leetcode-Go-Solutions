package main

import (
	"fmt"
	"sort"
)

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	a := kthLargest.Add(3)  // return 4
	b := kthLargest.Add(5)  // return 5
	c := kthLargest.Add(10) // return 5
	d := kthLargest.Add(9)  // return 8
	e := kthLargest.Add(4)  // return 8

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)
}

type KthLargest struct {
	k     int
	queue *[]int
}

// INCOMPLETE
func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{
		k:     k,
		queue: &[]int{},
	}

	s.AddBulk(nums)

	return s
}

func (kl *KthLargest) AddBulk(vals []int) {
	*kl.queue = append(*kl.queue, vals...)
	sort.Ints(*kl.queue)

}

func (kl *KthLargest) Add(val int) int {
	*kl.queue = append(*kl.queue, val)
	sort.Ints(*kl.queue)

	return (*kl.queue)[len(*kl.queue)-kl.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
