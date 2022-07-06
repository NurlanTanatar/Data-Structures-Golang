package main

import "fmt"

type linkedNode struct {
	data int
	next *linkedNode
}

type LinkedList struct {
	head   *linkedNode
	length int
}

func (l *LinkedList) Prepend(n *linkedNode) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l LinkedList) SearchWithValue(value int) bool {
	if l.length == 0 {
		return false
	}

	previousToSearch := l.head
	for previousToSearch.next.data != value {
		if previousToSearch.next.next == nil {
			return false
		}
		previousToSearch = previousToSearch.next
	}
	return true

}

func LinkedListExample() {
	myList := LinkedList{}
	node1, node2, node3, node4 := &linkedNode{data: 14}, &linkedNode{data: 24}, &linkedNode{data: 34}, &linkedNode{data: 44}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Prepend(node4)
	myList.PrintListData()
	myList.DeleteWithValue(44) // deleting head
	myList.DeleteWithValue(55) // deleting non-existent node
	myList.PrintListData()
	myList.DeleteWithValue(24) // deleting node from middle
	myList.PrintListData()
	fmt.Println(myList.SearchWithValue(14)) // searching by value
	emptyList := LinkedList{}
	emptyList.DeleteWithValue(4)              // deleting from empty list
	fmt.Println(emptyList.SearchWithValue(2)) // Search from empty list
}
