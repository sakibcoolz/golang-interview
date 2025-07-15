package searching

import (
	"fmt"
	"sort"
	"strings"
)

/*
Problem: Binary Search (LeetCode #704)
Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index.
Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Constraints:
- 1 <= nums.length <= 10^4
- -10^4 < nums[i], target < 10^4
- All the integers in nums are unique.
- nums is sorted in ascending order.
*/

// Approach 1: Iterative Binary Search
// Time Complexity: O(log n), Space Complexity: O(1)
func binarySearchIterative(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // Avoid overflow

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Approach 2: Recursive Binary Search
// Time Complexity: O(log n), Space Complexity: O(log n) due to recursion
func binarySearchRecursive(nums []int, target int) int {
	return binarySearchHelper(nums, target, 0, len(nums)-1)
}

func binarySearchHelper(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearchHelper(nums, target, mid+1, right)
	} else {
		return binarySearchHelper(nums, target, left, mid-1)
	}
}

// Approach 3: Using built-in sort.Search (for comparison)
// Time Complexity: O(log n), Space Complexity: O(1)
func binarySearchBuiltIn(nums []int, target int) int {
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	if index < len(nums) && nums[index] == target {
		return index
	}
	return -1
}

// Approach 4: Template for finding leftmost position
// Time Complexity: O(log n), Space Complexity: O(1)
func binarySearchLeftmost(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// Approach 5: Template for finding rightmost position
// Time Complexity: O(log n), Space Complexity: O(1)
func binarySearchRightmost(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	left-- // Adjust to get the rightmost position
	if left >= 0 && left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// Example usage and test cases
func RunBinarySearchExamples() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
		desc   string
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, "Example 1 - target exists"},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, "Example 2 - target doesn't exist"},
		{[]int{5}, 5, "Single element - found"},
		{[]int{5}, 3, "Single element - not found"},
		{[]int{1, 2, 3, 4, 5}, 1, "Target at beginning"},
		{[]int{1, 2, 3, 4, 5}, 5, "Target at end"},
		{[]int{1, 2, 3, 4, 5}, 3, "Target in middle"},
		{[]int{1, 3, 3, 3, 5}, 3, "Duplicate elements"},
		{[]int{}, 1, "Empty array"},
	}

	fmt.Println("Binary Search Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Iterative: Time O(log n), Space O(1)")
	fmt.Println("🔸 Recursive: Time O(log n), Space O(log n)")
	fmt.Println("🔸 Built-in: Time O(log n), Space O(1)")
	fmt.Println("🔸 All variants have same time complexity")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: nums = %v, target = %d\n", tc.nums, tc.target)

		if len(tc.nums) > 0 {
			result1 := binarySearchIterative(tc.nums, tc.target)
			result2 := binarySearchRecursive(tc.nums, tc.target)
			result3 := binarySearchBuiltIn(tc.nums, tc.target)
			result4 := binarySearchLeftmost(tc.nums, tc.target)
			result5 := binarySearchRightmost(tc.nums, tc.target)

			fmt.Printf("Iterative Result: %d\n", result1)
			fmt.Printf("Recursive Result: %d\n", result2)
			fmt.Printf("Built-in Result: %d\n", result3)
			fmt.Printf("Leftmost Result: %d\n", result4)
			fmt.Printf("Rightmost Result: %d\n", result5)
		} else {
			fmt.Printf("All methods return: -1 (empty array)\n")
		}
	}
}
