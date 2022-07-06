package main

import "fmt"

type MinHeap struct {
	slice []int
}

func (h *MinHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.minHeapifyUp(len(h.slice) - 1)

}

func (h *MinHeap) Extract() int {
	extracted := h.slice[0]

	if len(h.slice) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	h.slice[0] = h.slice[len(h.slice)-1] // taking out last element and putting it in the root
	h.slice = h.slice[:len(h.slice)-1]   // shrinking down to actual size of heap

	h.minHeapifyDown(0)

	return extracted
}

func (h *MinHeap) minHeapifyUp(i int) {
	for h.slice[parent(i)] > h.slice[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *MinHeap) minHeapifyDown(i int) {
	lastIndex := len(h.slice) - 1
	l, r := left(i), right(i)
	childToCompare := lastIndex
	for l <= lastIndex {
		if l == lastIndex { // if there is only one child we gonna take it otherwise accessing to right child will cause panic
			childToCompare = l
		} else if h.slice[l] < h.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.slice[i] > h.slice[childToCompare] {
			h.swap(i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}

	}

}

func (h *MinHeap) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func MinHeapExample() {
	m := MinHeap{}
	buildHeap := []int{1, 2, 3, 4, 9, 8, 7, 6, 5}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}
