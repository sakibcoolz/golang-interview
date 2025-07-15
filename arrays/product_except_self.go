package arrays

import (
	"fmt"
	"strings"
)

/*
Problem: Product of Array Except Self (LeetCode #238)
Given an integer array nums, return an array answer such that answer[i] is equal to the product
of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:
- 2 <= nums.length <= 10^5
- -30 <= nums[i] <= 30
- The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Follow up: Can you solve the problem in O(1) extra space complexity?
(The output array does not count as extra space for space complexity analysis.)
*/

// Approach 1: Left and Right products with extra arrays
// Time Complexity: O(n), Space Complexity: O(n)
func productExceptSelfWithArrays(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate left products
	leftProducts := make([]int, n)
	leftProducts[0] = 1
	for i := 1; i < n; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
	}

	// Calculate right products
	rightProducts := make([]int, n)
	rightProducts[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightProducts[i] = rightProducts[i+1] * nums[i+1]
	}

	// Multiply left and right products
	for i := 0; i < n; i++ {
		result[i] = leftProducts[i] * rightProducts[i]
	}

	return result
}

// Approach 2: Optimized with O(1) extra space
// Time Complexity: O(n), Space Complexity: O(1)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// First pass: calculate left products
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Second pass: multiply by right products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// Example usage and test cases
func RunProductExceptSelfExamples() {
	// Test cases
	testCases := []struct {
		nums []int
		desc string
	}{
		{[]int{1, 2, 3, 4}, "Example 1 - positive numbers"},
		{[]int{-1, 1, 0, -3, 3}, "Example 2 - with zero and negatives"},
		{[]int{2, 3}, "Minimum case - two elements"},
		{[]int{1, 0}, "With zero"},
		{[]int{-1, -2, -3}, "All negative numbers"},
	}

	fmt.Println("Product of Array Except Self Problem Solutions:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Two Arrays: Time O(n), Space O(n)")
	fmt.Println("🔸 Optimized: Time O(n), Space O(1)")
	fmt.Println("🔸 Note: Both approaches require O(n) for output array")
	fmt.Println(strings.Repeat("-", 50))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.nums)

		result1 := productExceptSelfWithArrays(tc.nums)
		result2 := productExceptSelf(tc.nums)

		fmt.Printf("With Arrays Result: %v\n", result1)
		fmt.Printf("Optimized Result: %v\n", result2)
	}
}
