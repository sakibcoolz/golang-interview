package linkedlists

import (
	"fmt"
	"strings"
)

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Problem: Reverse Linked List (LeetCode #206)
Given the head of a singly linked list, reverse the list, and return the new head.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []

Constraints:
- The number of nodes in the list is the range [0, 5000].
- -5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

// Approach 1: Iterative
// Time Complexity: O(n), Space Complexity: O(1)
func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// Approach 2: Recursive
// Time Complexity: O(n), Space Complexity: O(n) due to recursion stack
func reverseListRecursive(head *ListNode) *ListNode {
	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	// Recursively reverse the rest of the list
	newHead := reverseListRecursive(head.Next)

	// Reverse the current connection
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// Approach 3: Using stack (for demonstration)
// Time Complexity: O(n), Space Complexity: O(n)
func reverseListStack(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Push all nodes to stack
	var stack []*ListNode
	current := head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// Pop and reconnect
	newHead := stack[len(stack)-1]
	current = newHead

	for i := len(stack) - 2; i >= 0; i-- {
		current.Next = stack[i]
		current = current.Next
	}
	current.Next = nil

	return newHead
}

// Helper function to create linked list from slice
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

// Example usage and test cases
func RunReverseLinkedListExamples() {
	// Test cases
	testCases := []struct {
		vals []int
		desc string
	}{
		{[]int{1, 2, 3, 4, 5}, "Example 1 - five nodes"},
		{[]int{1, 2}, "Example 2 - two nodes"},
		{[]int{}, "Example 3 - empty list"},
		{[]int{1}, "Single node"},
		{[]int{1, 2, 3}, "Three nodes"},
	}

	fmt.Println("Reverse Linked List Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Iterative: Time O(n), Space O(1)")
	fmt.Println("🔸 Recursive: Time O(n), Space O(n) - call stack")
	fmt.Println("🔸 Stack: Time O(n), Space O(n) - extra storage")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.vals)

		// Create three copies for testing different approaches
		head1 := createLinkedList(tc.vals)
		head2 := createLinkedList(tc.vals)
		head3 := createLinkedList(tc.vals)

		result1 := reverseListIterative(head1)
		result2 := reverseListRecursive(head2)
		result3 := reverseListStack(head3)

		fmt.Printf("Iterative Result: %v\n", linkedListToSlice(result1))
		fmt.Printf("Recursive Result: %v\n", linkedListToSlice(result2))
		fmt.Printf("Stack Result: %v\n", linkedListToSlice(result3))
	}
}
