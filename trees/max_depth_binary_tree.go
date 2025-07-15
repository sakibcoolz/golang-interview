package trees

import (
	"fmt"
	"strings"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Problem: Maximum Depth of Binary Tree (LeetCode #104)
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3

Example 2:
Input: root = [1,null,2]
Output: 2

Constraints:
- The number of nodes in the tree is in the range [0, 10^4].
- -100 <= Node.val <= 100
*/

// Approach 1: Recursive (DFS)
// Time Complexity: O(n), Space Complexity: O(h) where h is height of tree
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

// Approach 2: Iterative (BFS)
// Time Complexity: O(n), Space Complexity: O(w) where w is max width of tree
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// Approach 3: Iterative (DFS with stack)
// Time Complexity: O(n), Space Complexity: O(h)
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Stack to store (node, depth) pairs
	type StackItem struct {
		node  *TreeNode
		depth int
	}

	stack := []StackItem{{root, 1}}
	maxDepth := 0

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node, depth := item.node, item.depth
		maxDepth = max(maxDepth, depth)

		if node.Left != nil {
			stack = append(stack, StackItem{node.Left, depth + 1})
		}
		if node.Right != nil {
			stack = append(stack, StackItem{node.Right, depth + 1})
		}
	}

	return maxDepth
}

// Helper function to find maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to create binary tree from level-order slice
func createBinaryTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Example usage and test cases
func RunMaxDepthExamples() {
	// Test cases
	testCases := []struct {
		vals []interface{}
		desc string
	}{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, "Example 1 - balanced tree"},
		{[]interface{}{1, nil, 2}, "Example 2 - right skewed"},
		{[]interface{}{}, "Empty tree"},
		{[]interface{}{1}, "Single node"},
		{[]interface{}{1, 2, 3, 4, 5}, "Complete binary tree"},
		{[]interface{}{1, 2, nil, 3, nil, 4, nil, 5}, "Left skewed tree"},
	}

	fmt.Println("Maximum Depth of Binary Tree Problem Solutions:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Recursive: Time O(n), Space O(h)")
	fmt.Println("🔸 BFS Queue: Time O(n), Space O(w) - max width")
	fmt.Println("🔸 DFS Stack: Time O(n), Space O(h) - max height")
	fmt.Println("🔸 Best case: O(log n), Worst case: O(n)")
	fmt.Println(strings.Repeat("-", 50))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.vals)

		root := createBinaryTree(tc.vals)

		result1 := maxDepthRecursive(root)
		result2 := maxDepthIterative(root)
		result3 := maxDepthDFS(root)

		fmt.Printf("Recursive Result: %d\n", result1)
		fmt.Printf("BFS Iterative Result: %d\n", result2)
		fmt.Printf("DFS Iterative Result: %d\n", result3)
	}
}
