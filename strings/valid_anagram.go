package strings

import (
	"fmt"
	"sort"
	"strings"
)

/*
Problem: Valid Anagram (LeetCode #242)
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
- 1 <= s.length, t.length <= 5 * 10^4
- s and t consist of lowercase English letters only.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

// Approach 1: Sort both strings
// Time Complexity: O(n log n), Space Complexity: O(1) if we don't count output space
func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert to slices and sort
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")

	sort.Strings(sChars)
	sort.Strings(tChars)

	return strings.Join(sChars, "") == strings.Join(tChars, "")
}

// Approach 2: Character frequency count
// Time Complexity: O(n), Space Complexity: O(1) for lowercase English letters
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Count frequency of each character
	charCount := make(map[rune]int)

	// Count characters in s
	for _, char := range s {
		charCount[char]++
	}

	// Subtract characters in t
	for _, char := range t {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	// Check if all counts are zero
	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

// Approach 3: Using array for lowercase English letters (most efficient)
// Time Complexity: O(n), Space Complexity: O(1)
func isAnagramArray(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Use array for lowercase English letters
	charCount := make([]int, 26)

	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
		charCount[t[i]-'a']--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

// Example usage and test cases
func RunValidAnagramExamples() {
	// Test cases
	testCases := []struct {
		s    string
		t    string
		desc string
	}{
		{"anagram", "nagaram", "Example 1 - valid anagram"},
		{"rat", "car", "Example 2 - not anagram"},
		{"listen", "silent", "Valid anagram"},
		{"hello", "bello", "Different characters"},
		{"a", "ab", "Different lengths"},
		{"", "", "Empty strings"},
		{"aab", "aba", "Same characters, different count"},
	}

	fmt.Println("Valid Anagram Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n⏱️  Time Complexity Analysis:")
	fmt.Println("🔸 Sorting: Time O(n log n), Space O(1)")
	fmt.Println("🔸 Hash Map: Time O(n), Space O(1) - max 26 chars")
	fmt.Println("🔸 Array: Time O(n), Space O(1) - fixed 26 chars")
	fmt.Println(strings.Repeat("-", 40))

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: s = \"%s\", t = \"%s\"\n", tc.s, tc.t)

		result1 := isAnagramSort(tc.s, tc.t)
		result2 := isAnagram(tc.s, tc.t)
		result3 := isAnagramArray(tc.s, tc.t)

		fmt.Printf("Sort Method: %t\n", result1)
		fmt.Printf("Hash Map Method: %t\n", result2)
		fmt.Printf("Array Method: %t\n", result3)
	}
}
