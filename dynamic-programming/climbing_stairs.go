package dp

import (
	"fmt"
	"strings"
)

/*
Problem: Climbing Stairs (LeetCode #70)
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
- 1 <= n <= 45
*/

// Approach 1: Recursive (Naive) - Exponential time
// Time Complexity: O(2^n), Space Complexity: O(n)
func climbStairsRecursive(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

// Approach 2: Memoization (Top-down DP)
// Time Complexity: O(n), Space Complexity: O(n)
func climbStairsMemo(n int) int {
	memo := make(map[int]int)
	return climbStairsHelper(n, memo)
}

func climbStairsHelper(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = climbStairsHelper(n-1, memo) + climbStairsHelper(n-2, memo)
	return memo[n]
}

// Approach 3: Tabulation (Bottom-up DP)
// Time Complexity: O(n), Space Complexity: O(n)
func climbStairsTabulation(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Approach 4: Space Optimized (Fibonacci-like)
// Time Complexity: O(n), Space Complexity: O(1)
func climbStairsOptimized(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// Approach 5: Matrix Exponentiation (Advanced)
// Time Complexity: O(log n), Space Complexity: O(1)
func climbStairsMatrix(n int) int {
	if n <= 2 {
		return n
	}

	// Matrix [[1,1], [1,0]]^(n-1) * [2,1] = [fib(n+1), fib(n)]
	result := matrixPower([][]int{{1, 1}, {1, 0}}, n-1)
	return result[0][0]*2 + result[0][1]*1
}

func matrixPower(matrix [][]int, n int) [][]int {
	if n == 1 {
		return matrix
	}

	if n%2 == 0 {
		half := matrixPower(matrix, n/2)
		return matrixMultiply(half, half)
	}

	return matrixMultiply(matrix, matrixPower(matrix, n-1))
}

func matrixMultiply(a, b [][]int) [][]int {
	return [][]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

// Example usage and test cases
func RunClimbingStairsExamples() {
	// Test cases
	testCases := []struct {
		n    int
		desc string
	}{
		{2, "Example 1 - 2 steps"},
		{3, "Example 2 - 3 steps"},
		{1, "Base case - 1 step"},
		{4, "4 steps"},
		{5, "5 steps"},
		{10, "10 steps"},
		{20, "20 steps (larger input)"},
	}

	fmt.Println("Climbing Stairs Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Recursive: Time O(2ⁿ), Space O(n)")
	fmt.Println("🔸 Memoization: Time O(n), Space O(n)")
	fmt.Println("🔸 Tabulation: Time O(n), Space O(n)")
	fmt.Println("🔸 Space Optimized: Time O(n), Space O(1)")
	fmt.Println("🔸 Matrix Power: Time O(log n), Space O(1)")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: n = %d\n", tc.n)

		// For small inputs, show recursive result too
		if tc.n <= 10 {
			result1 := climbStairsRecursive(tc.n)
			fmt.Printf("Recursive Result: %d\n", result1)
		}

		result2 := climbStairsMemo(tc.n)
		result3 := climbStairsTabulation(tc.n)
		result4 := climbStairsOptimized(tc.n)
		result5 := climbStairsMatrix(tc.n)

		fmt.Printf("Memoization Result: %d\n", result2)
		fmt.Printf("Tabulation Result: %d\n", result3)
		fmt.Printf("Space Optimized Result: %d\n", result4)
		fmt.Printf("Matrix Exponentiation Result: %d\n", result5)
	}
}
