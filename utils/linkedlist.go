package utils

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// CreateLinkedList creates a linked list from a slice of integers
func CreateLinkedList(vals []int) *ListNode {
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

// LinkedListToSlice converts a linked list to a slice
func LinkedListToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
