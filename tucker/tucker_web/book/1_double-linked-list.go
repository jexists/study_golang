package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

func main() {
	root := &Node[int]{nil, nil, 10}
	root.next = &Node[int]{nil, root, 20}
	root.next.next = &Node[int]{nil, root.next, 30}

	tail := root.next.next
	for n := tail; n != nil; n = n.prev {
		fmt.Printf("node val:%d\n", n.val)
	}
	// node val:30
	// node val:20
	// node val:10
}
