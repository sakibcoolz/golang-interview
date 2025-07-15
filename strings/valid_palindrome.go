package strings

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Problem: Valid Palindrome (LeetCode #125)
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters, it reads the same forward and backward.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
- 1 <= s.length <= 2 * 10^5
- s consists only of printable ASCII characters.
*/

// Approach 1: Clean string then check
// Time Complexity: O(n), Space Complexity: O(n)
func isPalindromeClean(s string) bool {
	// Clean the string
	var cleaned strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(unicode.ToLower(char))
		}
	}

	cleanStr := cleaned.String()

	// Check palindrome
	left, right := 0, len(cleanStr)-1
	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// Approach 2: Two pointers without extra space
// Time Complexity: O(n), Space Complexity: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric characters from left
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// Skip non-alphanumeric characters from right
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// Compare characters (case insensitive)
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

// Helper function to check if character is alphanumeric
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// Helper function to convert to lowercase
func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// Example usage and test cases
func RunValidPalindromeExamples() {
	// Test cases
	testCases := []struct {
		s    string
		desc string
	}{
		{"A man, a plan, a canal: Panama", "Example 1 - complex palindrome"},
		{"race a car", "Example 2 - not palindrome"},
		{" ", "Example 3 - empty after cleanup"},
		{"Madam", "Simple palindrome"},
		{"No 'x' in Nixon", "Another palindrome"},
		{"Was it a car or a cat I saw?", "Palindrome with punctuation"},
		{"hello", "Simple non-palindrome"},
		{"", "Empty string"},
		{"a", "Single character"},
	}

	fmt.Println("Valid Palindrome Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Clean & Check: Time O(n), Space O(n)")
	fmt.Println("🔸 Two Pointers: Time O(n), Space O(1)")
	fmt.Println("🔸 Both approaches scan string once")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: \"%s\"\n", tc.s)

		result1 := isPalindromeClean(tc.s)
		result2 := isPalindrome(tc.s)

		fmt.Printf("Clean String Method: %t\n", result1)
		fmt.Printf("Two Pointers Method: %t\n", result2)
	}
}
