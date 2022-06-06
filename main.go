package main

import (
	"fmt"
	"golang_sorts/Sorts"
)

func main() {
  	arr := []int{5, 1, 4, 2}

	BubbleArr := Sort.BubbleSort(arr)
	fmt.Printf("Bubble Sort: ")
	fmt.Println(BubbleArr)
	
	SelectionArr := Sort.SelectionSort(arr)
	fmt.Printf("Selection Sort: ")
	fmt.Println(SelectionArr)

	InsertionArr := Sort.InsertionSort(arr)
	fmt.Printf("Insertion Sort: ")
	fmt.Println(InsertionArr)

	

	
}