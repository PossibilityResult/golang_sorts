package Sort

// Bubble Sort is the simplest sorting algorithm.
// It consists of repeatedly swapping elements if they are not in the same order.

// Pseudocode
// consider array [5, 1, 4, 2]
// 1st pass: [5, 1, 4, 2] -> [1, 5, 4, 2] -> [1, 4, 5, 2] -> [1, 4, 2, 5]
// 2nd pass: [1, 4, 2, 5] -> [1, 4, 2, 5] -> [1, 2, 4, 5] -> [1, 2, 4, 5]

// Performance
// Worst Case: Time Complexity = O(n^2), Auxiliary Space = O(1)

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var sorted bool = false
	var temp int
	cpy := make([]int, len(arr))
	copy(cpy, arr)

	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if cpy[i] > cpy[i+1] {
				sorted = false
				temp = cpy[i]
				cpy[i] = cpy[i+1]
				cpy[i+1] = temp
			}
		}
	}
	return cpy
}

// Selection sort works by repeatedly finding the minimum element
// and placing it at the beginning.

// Pseudocode
// consider array [5, 1, 4, 2]
// 1st pass: [5, 1, 4, 2] -> [1, 5, 4, 2]
// 2nd pass: [1, 5, 4, 2] -> [1, 2, 4, 5]

// Performance
// Worst Case: Time Complexity = O(n^2), Auxiliary Space = O(1)

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var minInd int
	var temp int
	cpy := make([]int, len(arr))
	copy(cpy, arr)

	for i := 0; i < len(cpy); i++ {
		minInd = i
		for j := i; j < len(cpy); j++ {
			if cpy[j] < cpy[minInd] {
				minInd = j
			}
		}
		temp = cpy[i]
		cpy[i] = cpy[minInd]
		cpy[minInd] = temp
	}
	return cpy
}
