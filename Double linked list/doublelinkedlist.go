package main

import "fmt"

type Node struct {
	Prev *Node
	Data int
	Next *Node
}

func main() {
	fmt.Println("Double Linked List ")

	n1 := &Node{Data: 10}
	n2 := &Node{Data: 20}
	n3 := &Node{Data: 30}
	n4 := &Node{Data: 40}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n2.Prev = n1
	n3.Prev = n2
	n4.Prev = n3

	current := n1

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}

}
