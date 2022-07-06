package main

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

	if len(h.slice) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	h.slice[0] = h.slice[len(h.slice)-1] // taking out last element and putting it in the root
	h.slice = h.slice[:len(h.slice)-1]   // shrinking down to actual size of heap

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
	lastIndex := len(h.slice) - 1
	l, r := left(i), right(i)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex { // if there is only one child we gonna take it otherwise accessing to right child will cause panic
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

func (h *MaxHeap) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func MaxHeapExample() {
	m := MaxHeap{}
	buildHeap := []int{3, 4, 1, 9, 5, 4, 24, 123, 41, 49, 33}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
