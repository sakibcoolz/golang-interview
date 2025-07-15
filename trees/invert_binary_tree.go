package trees

import (
	"fmt"
	"strings"
)

/*
Problem: Invert Binary Tree (LeetCode #226)
Given the root of a binary tree, invert the tree, and return its root.

Example 1:
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

Constraints:
- The number of nodes in the tree is in the range [0, 100].
- -100 <= Node.val <= 100
*/

// Approach 1: Recursive
// Time Complexity: O(n), Space Complexity: O(h) where h is height of tree
func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right children
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert subtrees
	invertTreeRecursive(root.Left)
	invertTreeRecursive(root.Right)

	return root
}

// Approach 2: Iterative using queue (BFS)
// Time Complexity: O(n), Space Complexity: O(w) where w is max width of tree
func invertTreeIterativeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap left and right children
		node.Left, node.Right = node.Right, node.Left

		// Add children to queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// Approach 3: Iterative using stack (DFS)
// Time Complexity: O(n), Space Complexity: O(h)
func invertTreeIterativeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Swap left and right children
		node.Left, node.Right = node.Right, node.Left

		// Add children to stack
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return root
}

// Helper function to convert tree to level-order slice for visualization
func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			result = append(result, nil)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Helper function to copy a tree (for testing multiple approaches)
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

// Example usage and test cases
func RunInvertBinaryTreeExamples() {
	// Test cases
	testCases := []struct {
		vals []interface{}
		desc string
	}{
		{[]interface{}{4, 2, 7, 1, 3, 6, 9}, "Example 1 - full tree"},
		{[]interface{}{2, 1, 3}, "Example 2 - simple tree"},
		{[]interface{}{}, "Example 3 - empty tree"},
		{[]interface{}{1}, "Single node"},
		{[]interface{}{1, 2}, "Two nodes"},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, "Complete binary tree"},
	}

	fmt.Println("Invert Binary Tree Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Recursive: Time O(n), Space O(h)")
	fmt.Println("🔸 BFS Queue: Time O(n), Space O(w) - max width")
	fmt.Println("🔸 DFS Stack: Time O(n), Space O(h) - max height")
	fmt.Println("🔸 h = height, w = width, n = nodes")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.vals)

		// Create three copies for testing different approaches
		root1 := createBinaryTree(tc.vals)
		root2 := copyTree(root1)
		root3 := copyTree(root1)

		result1 := invertTreeRecursive(root1)
		result2 := invertTreeIterativeBFS(root2)
		result3 := invertTreeIterativeDFS(root3)

		fmt.Printf("Recursive Result: %v\n", treeToSlice(result1))
		fmt.Printf("BFS Iterative Result: %v\n", treeToSlice(result2))
		fmt.Printf("DFS Iterative Result: %v\n", treeToSlice(result3))
	}
}
