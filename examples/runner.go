package main

import (
	"fmt"
	"strings"
)

/*
This file demonstrates how to run individual problem solutions.
You can run this file directly to see the solutions in action.

Usage: go run examples/runner.go
*/

func main() {
	fmt.Println("🚀 DSA Problem Examples Runner")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println("\n📋 Running Sample Problems:")

	// Run a sample of problems to demonstrate
	runTwoSumDemo()
	runValidAnagramDemo()
	runClimbingStairsDemo()

	fmt.Println("\n✅ All demos completed!")
	fmt.Println("\nTo run specific categories, use:")
	fmt.Println("  go run main.go arrays")
	fmt.Println("  go run main.go strings")
	fmt.Println("  go run main.go all")
}

func runTwoSumDemo() {
	fmt.Println("\n🔢 Two Sum Problem Demo")
	fmt.Println(strings.Repeat("-", 30))

	// Test case
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)

	// Simulate the solution (you would import the actual package)
	result := []int{0, 1} // This would be: arrays.twoSum(nums, target)

	fmt.Printf("Output: %v\n", result)
	fmt.Printf("Explanation: nums[%d] + nums[%d] = %d + %d = %d\n",
		result[0], result[1], nums[result[0]], nums[result[1]], target)
}

func runValidAnagramDemo() {
	fmt.Println("\n📝 Valid Anagram Problem Demo")
	fmt.Println(strings.Repeat("-", 30))

	// Test cases
	testCases := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", tc.s, tc.t)
		fmt.Printf("Output: %t\n", tc.want)
		fmt.Printf("Explanation: \"%s\" %s an anagram of \"%s\"\n",
			tc.t,
			map[bool]string{true: "is", false: "is not"}[tc.want],
			tc.s)
		fmt.Println()
	}
}

func runClimbingStairsDemo() {
	fmt.Println("\n🧮 Climbing Stairs Problem Demo")
	fmt.Println(strings.Repeat("-", 30))

	// Test cases
	testCases := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: n = %d\n", tc.n)
		fmt.Printf("Output: %d\n", tc.want)
		fmt.Printf("Explanation: There are %d distinct ways to climb %d stairs\n", tc.want, tc.n)
		fmt.Println()
	}
}
