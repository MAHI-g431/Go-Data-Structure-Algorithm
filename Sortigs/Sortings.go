package main

import "fmt"

func main() {
	fmt.Println("Sorting Techniues⭐")
	fmt.Println("Sorting Menu...")
	for {
		fmt.Println("1. Bubble Sort.🫧")
		fmt.Println("2. Insertion Sort.😜")
		fmt.Println("3. Selection Sort.😉")
		fmt.Println("4. Exit.🔙")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			{
				fmt.Println(choice, "Bubble Sort...")
				Bubble_Sort()
			}
		case 2:
			{
				fmt.Println(choice, "Insertion  Sort...")
				Insertion_Sort()
			}
		case 3:
			{
				fmt.Println(choice, "Selection Sort...")
				Selection_Sort()
			}
		case 4:
			{
				return
			}
		default:
			fmt.Println("Please Enter Valid Number...")
		}
	}
}
func Bubble_Sort() {
	fmt.Println("------------------#Bubble Sort#------------------")
	fmt.Println("Enter The Number Of Elements In An Array:")
	var no_of_elements int
	fmt.Scan(&no_of_elements)

	Array := make([]int, no_of_elements)

	fmt.Println("Enter The Array Elements:")
	for i := 0; i < no_of_elements; i++ {
		fmt.Scan(&Array[i])
	}

	fmt.Println("Before Sorting:", Array)

	n := len(Array)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if Array[j] > Array[j+1] {
				Array[j], Array[j+1] = Array[j+1], Array[j]
			}

		}
	}
	fmt.Println("After Sorting:", Array)

}
func Insertion_Sort() {
	var n int
	fmt.Println("------------------#Insertion Sort#------------------")
	fmt.Println("Enter Number Of Elements:")
	fmt.Scan(&n)
	Array := make([]int, n)
	fmt.Print("Enter Elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&Array[i])
	}

	fmt.Println("Before Sorting The Array:", Array)
	for i := 1; i < n; i++ {
		value := Array[i]
		j := i - 1
		for j >= 0 && Array[j] > value {
			Array[j+1] = Array[j]
			j--
		}
		Array[j+1] = value
	}
	fmt.Println("After Sorting The  Array:", Array)
}
func Selection_Sort() {
	var n int
	fmt.Println("Enter the number of elements:")
	fmt.Scan(&n)
	Array := make([]int, n)
	fmt.Println("Enter the Elements :")
	for i := 0; i < n; i++ {
		fmt.Scan(&Array[i])
	}
	fmt.Println("Before Sorting:", Array)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if Array[j] < Array[minIndex] {
				minIndex = j
			}
		}
		Array[i], Array[minIndex] = Array[minIndex], Array[i]
	}
	fmt.Println("After Sorting Array :", Array)
}
