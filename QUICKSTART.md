# 🚀 Golang Interview DSA Project - Quick Start Guide

## 🎯 What You Have

A comprehensive **Data Structures and Algorithms** practice repository based on **NeetCode** problems, implemented in **Go**. Perfect for coding interview preparation!

## 📁 What's Included

### ✅ **Fully Implemented Categories:**

1. **🔢 Arrays & Hashing (4 problems)**
   - Two Sum (Hash Map vs Brute Force)
   - Contains Duplicate (Optimized approaches)
   - Group Anagrams (Sorting vs Character Counting)
   - Product of Array Except Self (Space optimized)

2. **📝 Strings (2 problems)**
   - Valid Anagram (Multiple approaches)
   - Valid Palindrome (Two pointers technique)

3. **🔗 Linked Lists (2 problems)**
   - Reverse Linked List (Iterative, Recursive, Stack)
   - Merge Two Sorted Lists (Multiple approaches)

4. **🌳 Trees (2 problems)**
   - Maximum Depth of Binary Tree (DFS, BFS)
   - Invert Binary Tree (Multiple approaches)

5. **📚 Stack & Queue (1 problem)**
   - Valid Parentheses (Stack implementation)

6. **🔍 Searching (1 problem)**
   - Binary Search (5 different approaches)

7. **🔄 Sorting (3 algorithms)**
   - Merge Sort, Quick Sort, Insertion Sort

8. **🧮 Dynamic Programming (1 problem)**
   - Climbing Stairs (5 approaches including Matrix Exponentiation)

## 🚀 Quick Start (3 Minutes)

### 1. **Run Everything:**
```bash
cd golang-interview
go run main.go all
```

### 2. **Run Specific Category:**
```bash
go run main.go arrays      # Array problems
go run main.go strings     # String problems  
go run main.go trees       # Tree problems
```

### 3. **Run Tests:**
```bash
go test ./arrays -v        # Test array solutions
go test ./... -v           # Test everything
```

### 4. **See Examples:**
```bash
go run examples/runner.go  # See problem demos
```

## 🎓 Learning Path

### **👶 Beginner (Start Here):**
1. `arrays/two_sum.go` - Classic hash map pattern
2. `arrays/contains_duplicate.go` - Set/hash table usage
3. `strings/valid_anagram.go` - Character counting
4. `strings/valid_palindrome.go` - Two pointers technique
5. `linkedlists/reverse_linked_list.go` - Pointer manipulation

### **🏃 Intermediate:**
1. `arrays/group_anagrams.go` - Grouping with keys
2. `arrays/product_except_self.go` - Space optimization
3. `linkedlists/merge_two_sorted_lists.go` - Merge technique
4. `trees/max_depth_binary_tree.go` - Tree traversal
5. `stack-queue/valid_parentheses.go` - Stack patterns

### **🚀 Advanced:**
1. `trees/invert_binary_tree.go` - Tree manipulation
2. `searching/binary_search.go` - Search templates
3. `sorting/sorting_algorithms.go` - Algorithm analysis
4. `dynamic-programming/climbing_stairs.go` - DP patterns

## 🔧 File Structure

```
golang-interview/
├── 📂 arrays/              # Array & hashing problems
├── 📂 strings/             # String manipulation  
├── 📂 linkedlists/         # Linked list problems
├── 📂 trees/               # Binary tree problems
├── 📂 stack-queue/         # Stack & queue problems
├── 📂 searching/           # Binary search problems
├── 📂 sorting/             # Sorting algorithms
├── 📂 dynamic-programming/ # DP problems
├── 📂 utils/               # Helper functions
├── 📂 examples/            # Demo runners
├── 📄 main.go              # Main menu
├── 📄 README.md            # Project overview
├── 📄 PROBLEMS.md          # Detailed problem list
└── 📄 SETUP.md             # Setup instructions
```

## 🎯 Key Features

### **🔄 Multiple Approaches**
Each problem includes multiple solution approaches:
- Brute force (for learning)
- Optimized solutions (for interviews)
- Space-optimized versions
- Different algorithmic techniques

### **📊 Complexity Analysis**
Every solution includes:
- Time complexity: O(?)
- Space complexity: O(?)  
- Trade-off explanations

### **🧪 Comprehensive Testing**
- Unit tests for all solutions
- Edge case coverage
- Performance benchmarks
- Example usage

### **📖 Educational Comments**
- Detailed problem descriptions
- Step-by-step explanations
- Common patterns highlighted
- Interview tips included

## 💡 Usage Examples

### **View a Problem:**
```go
// File: arrays/two_sum.go
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        if index, exists := numMap[complement]; exists {
            return []int{index, i}
        }
        numMap[num] = i
    }
    
    return []int{}
}
```

### **Run Problem Examples:**
Each file has a `RunXxxExamples()` function that demonstrates the solutions with test cases.

### **Check Implementation Status:**
See `PROBLEMS.md` for:
- ✅ Completed problems
- 🚧 High priority TODO
- 📝 Future implementations

## 🎯 Interview Prep Strategy

### **Week 1-2: Foundations**
- Master the basic patterns (hash maps, two pointers, etc.)
- Focus on arrays and strings problems
- Practice explaining your approach

### **Week 3-4: Data Structures**
- Linked lists and trees
- Stack and queue problems
- Binary search techniques

### **Week 5-6: Advanced**
- Dynamic programming
- Graph algorithms (coming soon)
- System design problems

### **Daily Practice:**
1. Pick 2-3 problems from different categories
2. Implement multiple approaches
3. Analyze time/space complexity
4. Practice explaining your solution

## 🔗 Resources

- **📺 NeetCode YouTube**: Video explanations for all problems
- **🌐 NeetCode.io**: Interactive practice platform
- **📚 LeetCode**: Original problem source
- **📖 This Repository**: Your implementation practice

## 🚀 Next Steps

1. **Explore the code** - Start with simple problems
2. **Run the examples** - See solutions in action  
3. **Add your solutions** - Practice implementing variants
4. **Study the patterns** - Understand the underlying techniques
5. **Take mock interviews** - Practice explaining your approach

## 📞 Need Help?

- Check `SETUP.md` for detailed setup instructions
- Look at `PROBLEMS.md` for the complete problem list
- Examine the test files for usage examples
- Review the comments in each solution file

---

**🎉 Happy Coding!** This repository gives you a solid foundation for technical interviews. Focus on understanding the patterns rather than memorizing solutions. The goal is to recognize similar problems and apply the right techniques.

**⭐ Pro Tip**: After solving each problem, try to explain it to someone else (or even to a rubber duck). If you can teach it, you truly understand it!
