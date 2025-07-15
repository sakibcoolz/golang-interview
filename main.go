package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang-interview/arrays"
	dp "golang-interview/dynamic-programming"
	"golang-interview/linkedlists"
	"golang-interview/searching"
	"golang-interview/sorting"
	stackqueue "golang-interview/stack-queue"
	stringpkg "golang-interview/strings"
	"golang-interview/trees"
)

func main() {
	fmt.Println("🚀 Golang Interview DSA Problems")
	fmt.Println("Based on NeetCode Practice Problems")
	fmt.Println(strings.Repeat("=", 50))

	if len(os.Args) < 2 {
		showMenu()
		return
	}

	category := os.Args[1]

	switch strings.ToLower(category) {
	case "arrays", "array":
		runArrayProblems()
	case "strings", "string":
		runStringProblems()
	case "linkedlists", "linkedlist", "ll":
		runLinkedListProblems()
	case "trees", "tree":
		runTreeProblems()
	case "stacks", "stack", "queues", "queue", "stackqueue":
		runStackQueueProblems()
	case "searching", "search", "binarysearch":
		runSearchingProblems()
	case "sorting", "sort":
		runSortingProblems()
	case "dp", "dynamic", "dynamicprogramming":
		runDynamicProgrammingProblems()
	case "all":
		runAllProblems()
	default:
		fmt.Printf("Unknown category: %s\n", category)
		showMenu()
	}
}

func showMenu() {
	fmt.Println("\nUsage: go run main.go <category>")
	fmt.Println("\nAvailable categories:")
	fmt.Println("  arrays          - Array and hashing problems")
	fmt.Println("  strings         - String manipulation problems")
	fmt.Println("  linkedlists     - Linked list problems")
	fmt.Println("  trees           - Binary tree problems")
	fmt.Println("  stackqueue      - Stack and queue problems")
	fmt.Println("  searching       - Binary search problems")
	fmt.Println("  sorting         - Sorting algorithms")
	fmt.Println("  dp              - Dynamic programming problems")
	fmt.Println("  all             - Run all categories")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go arrays")
	fmt.Println("  go run main.go strings")
	fmt.Println("  go run main.go all")
}

func runArrayProblems() {
	fmt.Println("\n🔢 ARRAY PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Two Sum Problem:")
	arrays.RunTwoSumExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("2. Group Anagrams Problem:")
	arrays.RunGroupAnagramsExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("3. Contains Duplicate Problem:")
	arrays.RunContainsDuplicateExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("4. Product of Array Except Self Problem:")
	arrays.RunProductExceptSelfExamples()

	fmt.Println("\n✅ Array problems completed!")
}

func runStringProblems() {
	fmt.Println("\n📝 STRING PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Valid Anagram Problem:")
	stringpkg.RunValidAnagramExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("2. Valid Palindrome Problem:")
	stringpkg.RunValidPalindromeExamples()

	fmt.Println("\n✅ String problems completed!")
}

func runLinkedListProblems() {
	fmt.Println("\n🔗 LINKED LIST PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Reverse Linked List Problem:")
	linkedlists.RunReverseLinkedListExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("2. Merge Two Sorted Lists Problem:")
	linkedlists.RunMergeTwoListsExamples()

	fmt.Println("\n✅ Linked list problems completed!")
}

func runTreeProblems() {
	fmt.Println("\n🌳 TREE PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Maximum Depth of Binary Tree Problem:")
	trees.RunMaxDepthExamples()

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("2. Invert Binary Tree Problem:")
	trees.RunInvertBinaryTreeExamples()

	fmt.Println("\n✅ Tree problems completed!")
}

func runStackQueueProblems() {
	fmt.Println("\n📚 STACK & QUEUE PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Valid Parentheses Problem:")
	stackqueue.RunValidParenthesesExamples()

	fmt.Println("\n✅ Stack & Queue problems completed!")
}

func runSearchingProblems() {
	fmt.Println("\n🔍 SEARCHING PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Binary Search Problem:")
	searching.RunBinarySearchExamples()

	fmt.Println("\n✅ Searching problems completed!")
}

func runSortingProblems() {
	fmt.Println("\n🔄 SORTING PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("Sorting Algorithms (Merge Sort, Quick Sort, Insertion Sort):")
	sorting.RunSortingExamples()

	fmt.Println("\n✅ Sorting problems completed!")
}

func runDynamicProgrammingProblems() {
	fmt.Println("\n🧮 DYNAMIC PROGRAMMING PROBLEMS")
	fmt.Println(strings.Repeat("-", 30))

	fmt.Println("1. Climbing Stairs Problem:")
	dp.RunClimbingStairsExamples()

	fmt.Println("\n✅ Dynamic Programming problems completed!")
}

func runAllProblems() {
	fmt.Println("\n🎯 RUNNING ALL PROBLEMS")
	fmt.Println(strings.Repeat("=", 50))

	runArrayProblems()
	runStringProblems()
	runLinkedListProblems()
	runTreeProblems()
	runStackQueueProblems()
	runSearchingProblems()
	runSortingProblems()
	runDynamicProgrammingProblems()

	fmt.Println("\n🎉 All problems completed!")
	fmt.Println("\nTo run individual problems, navigate to the specific files:")
	fmt.Println("  go run arrays/two_sum.go")
	fmt.Println("  go run strings/valid_anagram.go")
	fmt.Println("  go run linkedlists/reverse_linked_list.go")
	fmt.Println("  etc.")
}

// Helper function to parse integer from string
func parseInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}
