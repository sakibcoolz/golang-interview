package arrays

import (
	"fmt"
	"strings"
)

/*
Problem: Two Sum (LeetCode #1)
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
- 2 <= nums.length <= 10^4
- -10^9 <= nums[i] <= 10^9
- -10^9 <= target <= 10^9
- Only one valid answer exists.
*/

// Approach 1: Brute Force
// Time Complexity: O(n^2), Space Complexity: O(1)
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Approach 2: Hash Map (Optimal)
// Time Complexity: O(n), Space Complexity: O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, exists := numMap[complement]; exists {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return []int{}
}

// Example usage and test cases
func RunTwoSumExamples() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
		desc   string
	}{
		{[]int{2, 7, 11, 15}, 9, "Example 1"},
		{[]int{3, 2, 4}, 6, "Example 2"},
		{[]int{3, 3}, 6, "Example 3"},
		{[]int{-1, -2, -3, -4, -5}, -8, "Negative numbers"},
	}

	fmt.Println("Two Sum Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n📊 COMPLEXITY ANALYSIS:")
	fmt.Println("🔸 Brute Force: Time O(n²), Space O(1)")
	fmt.Println("🔸 Hash Map:    Time O(n),  Space O(n)")
	fmt.Println("🔸 Where n = length of input array")

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: nums = %v, target = %d\n", tc.nums, tc.target)

		result1 := twoSumBruteForce(tc.nums, tc.target)
		result2 := twoSum(tc.nums, tc.target)

		fmt.Printf("Brute Force Result [O(n²)]: %v\n", result1)
		fmt.Printf("Hash Map Result [O(n)]:     %v\n", result2)
	}
}
