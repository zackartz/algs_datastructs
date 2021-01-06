package datastructs

import "fmt"

type MaxHeap struct {
	slice []int
}

func (h *MaxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.slice[0]
	l := len(h.slice) - 1

	if len(h.slice) == 0 {
		fmt.Println("cannot extract because arr is empty")
		return -1
	}

	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.slice[parent(i)] < h.slice[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *MaxHeap) maxHeapifyDown(i int) {
	lastIdx := len(h.slice) - 1
	l, r := left(i), right(i)
	childToCompare := 0

	for l <= lastIdx {
		if l == lastIdx {
			childToCompare = l
		} else if h.slice[l] > h.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.slice[i] < h.slice[childToCompare] {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}
