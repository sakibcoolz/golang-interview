package arrays

import (
	"fmt"
	"sort"
	"strings"
)

/*
Problem: Group Anagrams (LeetCode #49)
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Constraints:
- 1 <= strs.length <= 10^4
- 0 <= strs[i].length <= 100
- strs[i] consists of lowercase English letters only.
*/

// Approach 1: Sort strings as keys
// Time Complexity: O(N * K log K) where N is length of strs and K is max length of string
// Space Complexity: O(N * K)
func groupAnagramsSort(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Sort the string to create a key
		chars := strings.Split(str, "")
		sort.Strings(chars)
		key := strings.Join(chars, "")

		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Approach 2: Character count as key
// Time Complexity: O(N * K) where N is length of strs and K is max length of string
// Space Complexity: O(N * K)
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		// Count characters
		count := make([]int, 26)
		for _, char := range str {
			count[char-'a']++
		}

		// Create key from character counts
		key := fmt.Sprintf("%v", count)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

// Example usage and test cases
func RunGroupAnagramsExamples() {
	// Test cases
	testCases := []struct {
		strs []string
		desc string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, "Example 1"},
		{[]string{""}, "Example 2"},
		{[]string{"a"}, "Example 3"},
		{[]string{"abc", "bca", "cab", "xyz", "zyx", "yxz"}, "Multiple groups"},
	}

	fmt.Println("Group Anagrams Problem Solutions:")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Println("\n📊 COMPLEXITY ANALYSIS:")
	fmt.Println("🔸 Sort Method:  Time O(N×K×log K), Space O(N×K)")
	fmt.Println("🔸 Count Method: Time O(N×K),       Space O(N×K)")
	fmt.Println("🔸 Where N = number of strings, K = maximum length of string")

	for _, tc := range testCases {
		fmt.Printf("\n%s:\n", tc.desc)
		fmt.Printf("Input: %v\n", tc.strs)

		result1 := groupAnagramsSort(tc.strs)
		result2 := groupAnagrams(tc.strs)

		fmt.Printf("Sort Method Result [O(N×K×log K)]: %v\n", result1)
		fmt.Printf("Count Method Result [O(N×K)]:      %v\n", result2)
	}
}
