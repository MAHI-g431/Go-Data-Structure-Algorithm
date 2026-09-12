package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func main() {

	var Head *Node

	for {

		fmt.Println("\n--- Linked List Menu ---")
		fmt.Println("1. Insert at beginning")
		fmt.Println("2. Insert at end")
		fmt.Println("3. Display")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			// Insert at beginning
			var n int
			fmt.Print("Enter Total number of values to be inserted: ")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				var value int
				fmt.Println("Enter Value:")
				fmt.Scan(&value)
				newnode := &Node{Data: value}

				if Head == nil {
					Head = newnode
				} else {
					current := Head
					for current.Next != nil {
						current = current.Next
					}
					current.Next = newnode
				}

				fmt.Println("Node inserted at beginning.")
			}

		case 2:
			// Insert at end

			var value int
			fmt.Print("Enter value: ")
			fmt.Scan(&value)

			newnode := &Node{Data: value}

			if Head == nil {
				Head = newnode
			} else {
				current := Head

				for current.Next != nil {
					current = current.Next
				}

				current.Next = newnode
			}

			fmt.Println("Node inserted at end.")

		case 3:
			// Display

			if Head == nil {
				fmt.Println("List is empty.")
			} else {
				current := Head

				fmt.Println("Linked List:")

				for current != nil {
					fmt.Print(current.Data, " ")
					current = current.Next
				}

				fmt.Println()
			}

		case 4:
			fmt.Println("Program exited.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
