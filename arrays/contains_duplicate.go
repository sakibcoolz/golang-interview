package arrays

import (
	"fmt"
	"strings"
)

/*
Problem: Contains Duplicate (LeetCode #217)
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
- 1 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
*/

// Approach 1: Hash Set
// Time Complexity: O(n), Space Complexity: O(n)
func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

// Approach 2: Using map with early return optimization
// Time Complexity: O(n), Space Complexity: O(n)
func containsDuplicateOptimized(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}

// Example usage and test cases
func RunContainsDuplicateExamples() {
	// Test cases
	testCases := []struct {
		nums []int
		desc string
	}{
		{[]int{1, 2, 3, 1}, "Example 1 - has duplicate"},
		{[]int{1, 2, 3, 4}, "Example 2 - no duplicates"},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, "Example 3 - multiple duplicates"},
		{[]int{1}, "Single element"},
		{[]int{}, "Empty array"},
	}

	fmt.Println("Contains Duplicate Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 HashSet: Time O(n), Space O(n)")
	fmt.Println("🔸 Sorting: Time O(n log n), Space O(1)")
	fmt.Println("🔸 Brute Force: Time O(n²), Space O(1)")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.nums)

		result1 := containsDuplicate(tc.nums)
		result2 := containsDuplicateOptimized(tc.nums)

		fmt.Printf("Hash Set Result: %t\n", result1)
		fmt.Printf("Optimized Result: %t\n", result2)
	}
}
