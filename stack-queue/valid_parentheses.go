package stackqueue

import (
	"fmt"
	"strings"
)

/*
Problem: Valid Parentheses (LeetCode #20)
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
- 1 <= s.length <= 10^4
- s consists of parentheses only '()[]{}'.
*/

// Approach 1: Using slice as stack
// Time Complexity: O(n), Space Complexity: O(n)
func isValidParentheses(s string) bool {
	// Map of closing to opening brackets
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack []rune

	for _, char := range s {
		// If it's a closing bracket
		if opening, exists := pairs[char]; exists {
			// Check if stack is empty or top doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	// Valid if stack is empty
	return len(stack) == 0
}

// Approach 2: Using custom stack structure
// Time Complexity: O(n), Space Complexity: O(n)
type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func isValidParenthesesWithStack(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := Stack{}

	for _, char := range s {
		if opening, exists := pairs[char]; exists {
			// It's a closing bracket
			if top, ok := stack.Pop(); !ok || top != opening {
				return false
			}
		} else {
			// It's an opening bracket
			stack.Push(char)
		}
	}

	return stack.IsEmpty()
}

// Approach 3: Counter-based (works only for single type, shown for comparison)
// This approach doesn't work for mixed types but is educational
func isValidSingleType(s string) bool {
	count := 0
	for _, char := range s {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

// Example usage and test cases
func RunValidParenthesesExamples() {
	// Test cases
	testCases := []struct {
		s    string
		desc string
	}{
		{"()", "Example 1 - simple valid"},
		{"()[]{}", "Example 2 - multiple types valid"},
		{"(]", "Example 3 - invalid"},
		{"([)]", "Interleaved - invalid"},
		{"{[]}", "Nested - valid"},
		{"((", "Unmatched opening"},
		{"))", "Unmatched closing"},
		{"", "Empty string"},
		{"(((())))", "Deeply nested valid"},
		{"([{}])", "Complex nested valid"},
	}

	fmt.Println("Valid Parentheses Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Slice Stack: Time O(n), Space O(n)")
	fmt.Println("🔸 Custom Stack: Time O(n), Space O(n)")
	fmt.Println("🔸 Counter (single type): Time O(n), Space O(1)")
	fmt.Println("🔸 All traverse string once")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: \"%s\"\n", tc.s)

		result1 := isValidParentheses(tc.s)
		result2 := isValidParenthesesWithStack(tc.s)

		fmt.Printf("Slice Stack Result: %t\n", result1)
		fmt.Printf("Custom Stack Result: %t\n", result2)

		// For demonstration of single type (only for strings with just one bracket type)
		if strings.ContainsAny(tc.s, "()") && !strings.ContainsAny(tc.s, "{}[]") {
			result3 := isValidSingleType(tc.s)
			fmt.Printf("Single Type Counter Result: %t\n", result3)
		}
	}
}
