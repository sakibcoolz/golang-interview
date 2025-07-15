# Problem List - NeetCode DSA Practice

This document lists all the implemented problems organized by category, based on NeetCode practice problems.

## 📊 Progress Overview

- ✅ **Arrays & Hashing**: 4/15 problems
- ✅ **Strings**: 2/10 problems  
- ✅ **Linked Lists**: 2/8 problems
- ✅ **Trees**: 2/12 problems
- ✅ **Stack & Queue**: 1/8 problems
- ✅ **Searching**: 1/5 problems
- ✅ **Sorting**: 3/3 problems
- ✅ **Dynamic Programming**: 1/10 problems

## 🔢 Arrays & Hashing

### ✅ Implemented
1. **Two Sum** - `arrays/two_sum.go`
   - Problem: Find indices of two numbers that add up to target
   - Approaches: Brute force O(n²), Hash map O(n)
   - LeetCode: #1

2. **Contains Duplicate** - `arrays/contains_duplicate.go`
   - Problem: Check if array contains duplicates
   - Approaches: Hash set O(n)
   - LeetCode: #217

3. **Group Anagrams** - `arrays/group_anagrams.go`
   - Problem: Group strings that are anagrams
   - Approaches: Sort keys, Character counting
   - LeetCode: #49

4. **Product of Array Except Self** - `arrays/product_except_self.go`
   - Problem: Product of all elements except self
   - Approaches: Left/right products, Space optimized
   - LeetCode: #238

### 🚧 TODO - High Priority
5. **Top K Frequent Elements** - LeetCode #347
6. **Valid Sudoku** - LeetCode #36
7. **Longest Consecutive Sequence** - LeetCode #128
8. **Encode and Decode Strings** - LeetCode #271

### 📝 TODO - Medium Priority
9. **3Sum** - LeetCode #15
10. **Container With Most Water** - LeetCode #11
11. **Trapping Rain Water** - LeetCode #42

## 📝 Strings

### ✅ Implemented
1. **Valid Anagram** - `strings/valid_anagram.go`
   - Problem: Check if two strings are anagrams
   - Approaches: Sorting, Character counting, Array counting
   - LeetCode: #242

2. **Valid Palindrome** - `strings/valid_palindrome.go`
   - Problem: Check if string is palindrome (alphanumeric only)
   - Approaches: Clean then check, Two pointers
   - LeetCode: #125

### 🚧 TODO - High Priority
3. **Longest Substring Without Repeating Characters** - LeetCode #3
4. **Longest Repeating Character Replacement** - LeetCode #424
5. **Minimum Window Substring** - LeetCode #76

### 📝 TODO - Medium Priority
6. **String to Integer (atoi)** - LeetCode #8
7. **Find All Anagrams in a String** - LeetCode #438

## 🔗 Linked Lists

### ✅ Implemented
1. **Reverse Linked List** - `linkedlists/reverse_linked_list.go`
   - Problem: Reverse a singly linked list
   - Approaches: Iterative, Recursive, Stack
   - LeetCode: #206

2. **Merge Two Sorted Lists** - `linkedlists/merge_two_sorted_lists.go`
   - Problem: Merge two sorted linked lists
   - Approaches: Iterative, Recursive, In-place
   - LeetCode: #21

### 🚧 TODO - High Priority
3. **Linked List Cycle** - LeetCode #141
4. **Remove Nth Node From End** - LeetCode #19
5. **Reorder List** - LeetCode #143

### 📝 TODO - Medium Priority
6. **Add Two Numbers** - LeetCode #2
7. **Copy List with Random Pointer** - LeetCode #138

## 🌳 Trees

### ✅ Implemented
1. **Maximum Depth of Binary Tree** - `trees/max_depth_binary_tree.go`
   - Problem: Find maximum depth of binary tree
   - Approaches: Recursive DFS, Iterative BFS, Iterative DFS
   - LeetCode: #104

2. **Invert Binary Tree** - `trees/invert_binary_tree.go`
   - Problem: Invert/flip a binary tree
   - Approaches: Recursive, Iterative BFS, Iterative DFS
   - LeetCode: #226

### 🚧 TODO - High Priority
3. **Same Tree** - LeetCode #100
4. **Subtree of Another Tree** - LeetCode #572
5. **Lowest Common Ancestor of BST** - LeetCode #235
6. **Binary Tree Level Order Traversal** - LeetCode #102
7. **Validate Binary Search Tree** - LeetCode #98

### 📝 TODO - Medium Priority
8. **Kth Smallest Element in BST** - LeetCode #230
9. **Construct Binary Tree from Preorder and Inorder** - LeetCode #105

## 📚 Stack & Queue

### ✅ Implemented
1. **Valid Parentheses** - `stack-queue/valid_parentheses.go`
   - Problem: Check if parentheses are valid
   - Approaches: Stack using slice, Custom stack
   - LeetCode: #20

### 🚧 TODO - High Priority
2. **Min Stack** - LeetCode #155
3. **Evaluate Reverse Polish Notation** - LeetCode #150
4. **Generate Parentheses** - LeetCode #22
5. **Daily Temperatures** - LeetCode #739

## 🔍 Searching

### ✅ Implemented
1. **Binary Search** - `searching/binary_search.go`
   - Problem: Search target in sorted array
   - Approaches: Iterative, Recursive, Built-in, Templates
   - LeetCode: #704

### 🚧 TODO - High Priority
2. **Search in Rotated Sorted Array** - LeetCode #33
3. **Find Minimum in Rotated Sorted Array** - LeetCode #153
4. **Search a 2D Matrix** - LeetCode #74
5. **Time Based Key-Value Store** - LeetCode #981

## 🔄 Sorting

### ✅ Implemented
1. **Merge Sort** - `sorting/sorting_algorithms.go`
   - Time: O(n log n), Space: O(n), Stable
   - Divide and conquer approach

2. **Quick Sort** - `sorting/sorting_algorithms.go`
   - Time: O(n log n) avg, O(n²) worst, Space: O(log n)
   - Pivot-based partitioning

3. **Insertion Sort** - `sorting/sorting_algorithms.go`
   - Time: O(n²) worst, O(n) best, Space: O(1), Stable
   - Good for small arrays

### 📝 TODO - Additional Algorithms
4. **Heap Sort** - O(n log n), O(1) space
5. **Counting Sort** - O(n + k), good for limited range
6. **Radix Sort** - O(d × (n + k)), for integers

## 🧮 Dynamic Programming

### ✅ Implemented
1. **Climbing Stairs** - `dynamic-programming/climbing_stairs.go`
   - Problem: Count ways to climb n stairs
   - Approaches: Recursive, Memoization, Tabulation, Space optimized, Matrix
   - LeetCode: #70

### 🚧 TODO - High Priority
2. **House Robber** - LeetCode #198
3. **Coin Change** - LeetCode #322
4. **Longest Increasing Subsequence** - LeetCode #300
5. **Word Break** - LeetCode #139

### 📝 TODO - Medium Priority
6. **Combination Sum IV** - LeetCode #377
7. **House Robber II** - LeetCode #213
8. **Decode Ways** - LeetCode #91
9. **Unique Paths** - LeetCode #62
10. **Jump Game** - LeetCode #55

## 📈 Graphs (Future Implementation)

### 🚧 TODO - High Priority
1. **Number of Islands** - LeetCode #200
2. **Clone Graph** - LeetCode #133
3. **Pacific Atlantic Water Flow** - LeetCode #417
4. **Course Schedule** - LeetCode #207
5. **Graph Valid Tree** - LeetCode #261

## 🏗️ Design Problems (Future Implementation)

### 🚧 TODO
1. **Design Data Structures**
   - Dynamic Array (Resizable Array)
   - Singly Linked List
   - Hash Table
   - Binary Search Tree
   - Min/Max Heap

## 🎯 Next Implementation Priority

1. **Arrays**: Top K Frequent Elements, Valid Sudoku
2. **Strings**: Longest Substring Without Repeating Characters
3. **Linked Lists**: Linked List Cycle, Remove Nth Node
4. **Trees**: Same Tree, Binary Tree Level Order Traversal
5. **Stack/Queue**: Min Stack, Generate Parentheses
6. **DP**: House Robber, Coin Change

## 📚 Learning Path Suggestion

### Beginner (Start Here)
1. Two Sum
2. Contains Duplicate
3. Valid Anagram
4. Valid Palindrome
5. Reverse Linked List
6. Maximum Depth of Binary Tree
7. Valid Parentheses
8. Binary Search
9. Climbing Stairs

### Intermediate
1. Group Anagrams
2. Product of Array Except Self
3. Merge Two Sorted Lists
4. Invert Binary Tree
5. Longest Substring Without Repeating Characters
6. Linked List Cycle
7. Min Stack
8. House Robber

### Advanced
1. Longest Consecutive Sequence
2. Minimum Window Substring
3. Reorder List
4. Binary Tree Level Order Traversal
5. Generate Parentheses
6. Search in Rotated Sorted Array
7. Coin Change
8. Course Schedule

---

**Note**: This is a living document. Problems are continuously being added and improved. Each problem includes multiple solution approaches with different time/space complexities to help understand various optimization techniques.
