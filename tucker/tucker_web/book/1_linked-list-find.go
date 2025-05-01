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
	node := root
	for i := 0; i < 9; i++ {
		node = Append(node, &Node[int]{nil, (i + 2) * 10})
	}
	for n := root; n != nil; n = n.next {
		fmt.Println("val: ", n.val)
	}
	// Array에서는 찾기
	var a [100]int
	fmt.Println("7th Array Val: ", a[6])
	// 7th Array Val:  0

	// Node에서 찾기
	node = root
	for i := 0; i < 6; i++ {
		node = node.next
	}
	fmt.Println("7th Node Val: ", node)
	// 7th Node Val:  &{0xc0000142c0 70}
}
