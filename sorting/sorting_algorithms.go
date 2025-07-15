package sorting

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Problem: Merge Sort
Merge Sort is a divide-and-conquer algorithm that divides the input array into two halves,
sorts them recursively, and then merges the sorted halves.

Time Complexity: O(n log n) in all cases
Space Complexity: O(n)
Stability: Stable
*/

// Approach 1: Standard Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying original
	result := make([]int, len(arr))
	copy(result, arr)

	mergeSortHelper(result, 0, len(result)-1)
	return result
}

func mergeSortHelper(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	// Sort left and right halves
	mergeSortHelper(arr, left, mid)
	mergeSortHelper(arr, mid+1, right)

	// Merge sorted halves
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	// Create temporary arrays
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)

	// Copy data to temp arrays
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	// Merge temp arrays back
	i, j, k := 0, 0, left

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

/*
Problem: Quick Sort
Quick Sort is a divide-and-conquer algorithm that picks a 'pivot' element and partitions
the array around the pivot.

Average Time Complexity: O(n log n)
Worst Time Complexity: O(n²)
Space Complexity: O(log n) for recursion
Stability: Not stable
*/

// Approach 1: Standard Quick Sort (Last element as pivot)
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Make a copy to avoid modifying original
	result := make([]int, len(arr))
	copy(result, arr)

	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Partition and get pivot index
		pi := partition(arr, low, high)

		// Sort elements before and after partition
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose last element as pivot
	pivot := arr[high]
	i := low - 1 // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Approach 2: Quick Sort with random pivot (better average case)
func quickSortRandom(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)

	quickSortRandomHelper(result, 0, len(result)-1)
	return result
}

func quickSortRandomHelper(arr []int, low, high int) {
	if low < high {
		// Random pivot
		randomIndex := low + rand.Intn(high-low+1)
		arr[randomIndex], arr[high] = arr[high], arr[randomIndex]

		pi := partition(arr, low, high)
		quickSortRandomHelper(arr, low, pi-1)
		quickSortRandomHelper(arr, pi+1, high)
	}
}

/*
Problem: Insertion Sort
Simple sorting algorithm that builds the final sorted array one item at a time.

Time Complexity: O(n²) worst case, O(n) best case
Space Complexity: O(1)
Stability: Stable
*/

func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// Example usage and test cases
func RunSortingExamples() {
	// Test cases
	testCases := []struct {
		arr  []int
		desc string
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, "Random array"},
		{[]int{5, 4, 3, 2, 1}, "Reverse sorted"},
		{[]int{1, 2, 3, 4, 5}, "Already sorted"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
		{[]int{3, 3, 3, 3}, "All same elements"},
		{[]int{1, 3, 2}, "Small array"},
	}

	fmt.Println("Sorting Algorithms Problem Solutions:")
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Merge Sort: Time O(n log n), Space O(n)")
	fmt.Println("🔸 Quick Sort: Time O(n log n) avg, O(n²) worst, Space O(log n)")
	fmt.Println("🔸 Quick Sort Random: Time O(n log n) avg, Space O(log n)")
	fmt.Println("🔸 Insertion Sort: Time O(n²), Space O(1)")
	fmt.Println(strings.Repeat("-", 45))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.arr)

		if len(tc.arr) > 0 {
			result1 := mergeSort(tc.arr)
			result2 := quickSort(tc.arr)
			result3 := quickSortRandom(tc.arr)
			result4 := insertionSort(tc.arr)

			fmt.Printf("Merge Sort Result: %v\n", result1)
			fmt.Printf("Quick Sort Result: %v\n", result2)
			fmt.Printf("Quick Sort (Random) Result: %v\n", result3)
			fmt.Printf("Insertion Sort Result: %v\n", result4)
		} else {
			fmt.Printf("All algorithms return: [] (empty array)\n")
		}
	}
}
