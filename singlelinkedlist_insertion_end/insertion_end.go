package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func main() {
	n1 := &Node{Data: 10}
	n2 := &Node{Data: 20}
	n3 := &Node{Data: 30}

	n1.Next = n2
	n2.Next = n3

	head := n1

	newnode := &Node{Data: 40}
	current := head

	for current.Next != nil {
		current = current.Next
	}
	current.Next = newnode

	current = head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}

}
