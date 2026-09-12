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

	Head := n1

	newnode := &Node{Data: 5}
	newnode.Next = Head
	Head = newnode

	current := Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}

}
