package main

import "fmt"

type treeNode struct {
	Key     int
	Counter int
	Left    *treeNode
	Right   *treeNode
}

func (n *treeNode) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &treeNode{Key: k, Counter: 1}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &treeNode{Key: k, Counter: 1}
		} else {
			n.Left.Insert(k)
		}
	} else {
		n.Counter++
	}
}

func (n *treeNode) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func (n *treeNode) InorderPrint() {
	if n == nil {
		return
	}
	n.Left.InorderPrint()
	fmt.Printf("%d(%d) ", n.Key, n.Counter)
	n.Right.InorderPrint()
}

func BinarySearchTreeExample() {
	tree := &treeNode{Key: 100, Counter: 1}
	tree.Insert(10)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(234)
	tree.Insert(43)
	tree.Insert(14)
	tree.Insert(86)
	fmt.Println(tree.Search(234))
	fmt.Println(tree.Search(235))
	tree.InorderPrint()
}
