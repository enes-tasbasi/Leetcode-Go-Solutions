package utils

type MaxHeap struct {
	heap []int
}

func (h *MaxHeap) Push(node int) {
	i := len(h.heap)
	h.heap = append(h.heap, node)
	j := i / 2
	for h.heap[j] < h.heap[i] {
		h.heap[j], h.heap[i], i, j = h.heap[i], h.heap[j], j, j/2
	}
}

func (h *MaxHeap) Pop() int {
	var node int
	if len(h.heap) > 0 {
		node = h.heap[0]

		h.heap[0] = h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]

		var i, il, ir int
		var ml bool
		var mr bool

		for {

			il = i*2 + 1
			ir = i*2 + 2
			ml = len(h.heap) > il
			mr = len(h.heap) > ir

			if ml {
				if h.heap[il] > h.heap[i] {
					if mr {
						if h.heap[il] >= h.heap[ir] {
							h.heap[il], h.heap[i], i = h.heap[i], h.heap[il], il
							continue
						}
					} else {
						h.heap[il], h.heap[i], i = h.heap[i], h.heap[il], il
						continue
					}
				}
			}

			if mr {
				if h.heap[ir] > h.heap[i] {
					if ml {
						if h.heap[ir] >= h.heap[il] {
							h.heap[ir], h.heap[i], i = h.heap[i], h.heap[ir], ir
							continue
						}
					} else {
						h.heap[ir], h.heap[i], i = h.heap[i], h.heap[ir], ir
						continue
					}
				}
			}

			return node
		}
	}

	return node
}
