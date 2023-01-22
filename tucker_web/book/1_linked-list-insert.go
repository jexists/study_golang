package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

func main() {
	root := &Node[int]{nil, 10}
	// node := Append(root, &Node[int]{nil, 20})
	node := root
	for i := 0; i < 3; i++ {
		node = Append(node, &Node[int]{nil, (i + 2) * 10})
	}
	for n := root; n != nil; n = n.next {
		fmt.Println("val: ", n.val)
	}
	// val:  10
	// val:  20
	// val:  30
	// val:  40

	fmt.Println("========")

	node = root.next          // 20 node
	originalNext := node.next // 30 node
	node = Append(node, &Node[int]{nil, 100})
	node.next = originalNext

	for n := root; n != nil; n = n.next {
		fmt.Println("val: ", n.val)
	}
	// val:  10
	// val:  20
	// val:  100
	// val:  30
	// val:  40
}
