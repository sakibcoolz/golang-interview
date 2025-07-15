package linkedlists

import (
	"fmt"
	"strings"
)

/*
Problem: Merge Two Sorted Lists (LeetCode #21)
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together
the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

Constraints:
- The number of nodes in both lists is in the range [0, 50].
- -100 <= Node.val <= 100
- Both list1 and list2 are sorted in non-decreasing order.
*/

// Approach 1: Iterative with dummy node
// Time Complexity: O(m + n), Space Complexity: O(1)
func mergeTwoListsIterative(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify the logic
	dummy := &ListNode{}
	current := dummy

	// Merge while both lists have nodes
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append remaining nodes
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// Approach 2: Recursive
// Time Complexity: O(m + n), Space Complexity: O(m + n) due to recursion stack
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	// Base cases
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Recursive case
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}

// Approach 3: In-place merging (more complex but no extra nodes)
// Time Complexity: O(m + n), Space Complexity: O(1)
func mergeTwoListsInPlace(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Ensure list1 starts with the smaller value
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}

	head := list1

	for list1.Next != nil && list2 != nil {
		if list1.Next.Val <= list2.Val {
			list1 = list1.Next
		} else {
			temp := list1.Next
			list1.Next = list2
			list2 = temp
			list1 = list1.Next
		}
	}

	if list2 != nil {
		list1.Next = list2
	}

	return head
}

// Example usage and test cases
func RunMergeTwoListsExamples() {
	// Test cases
	testCases := []struct {
		list1 []int
		list2 []int
		desc  string
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, "Example 1 - both non-empty"},
		{[]int{}, []int{}, "Example 2 - both empty"},
		{[]int{}, []int{0}, "Example 3 - one empty"},
		{[]int{1, 3, 5}, []int{2, 4, 6}, "Alternating elements"},
		{[]int{1, 1, 1}, []int{2, 2, 2}, "Same values in each list"},
		{[]int{1}, []int{2}, "Single nodes"},
	}

	fmt.Println("Merge Two Sorted Lists Problem Solutions:")
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Dummy Node: Time O(m+n), Space O(1)")
	fmt.Println("🔸 Recursive: Time O(m+n), Space O(m+n) - call stack")
	fmt.Println("🔸 In-place: Time O(m+n), Space O(1)")
	fmt.Println(strings.Repeat("-", 45))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: list1 = %v, list2 = %v\n", tc.list1, tc.list2)

		// Create three copies for testing different approaches
		l1_iter := createLinkedList(tc.list1)
		l2_iter := createLinkedList(tc.list2)
		l1_rec := createLinkedList(tc.list1)
		l2_rec := createLinkedList(tc.list2)
		l1_place := createLinkedList(tc.list1)
		l2_place := createLinkedList(tc.list2)

		result1 := mergeTwoListsIterative(l1_iter, l2_iter)
		result2 := mergeTwoListsRecursive(l1_rec, l2_rec)
		result3 := mergeTwoListsInPlace(l1_place, l2_place)

		fmt.Printf("Iterative Result: %v\n", linkedListToSlice(result1))
		fmt.Printf("Recursive Result: %v\n", linkedListToSlice(result2))
		fmt.Printf("In-place Result: %v\n", linkedListToSlice(result3))
	}
}
