package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	// 맨 처음 Node (ROOT)를 알고 있어야한다

	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d \n", node.val)
	// 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
}

// 추가하기
func AddNode(root *Node, val int) {
	var tail *Node
	tail = root
	// 맨 끝 찾기
	for tail.next != nil {
		tail = tail.next
	}

	// 새로운 값 추가
	node := &Node{val: val}
	tail.next = node
}
