# Setup Instructions for Golang Interview DSA Project

## Prerequisites

1. **Go Installation**: Make sure you have Go 1.21 or later installed
   ```bash
   go version
   ```

2. **Git** (optional): For version control
   ```bash
   git --version
   ```

## Project Setup

### 1. Initialize the Project

Navigate to the project directory and initialize the Go module:

```bash
cd golang-interview
go mod init golang-interview
go mod tidy
```

### 2. Install Testing Dependencies

```bash
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/require
```

### 3. Verify Setup

Run a simple test to verify everything is working:

```bash
go test ./arrays -v
```

## Running the Project

### Option 1: Run Main Menu
```bash
go run main.go all          # Run all categories
go run main.go arrays       # Run array problems only
go run main.go strings      # Run string problems only
go run main.go linkedlists  # Run linked list problems
go run main.go trees        # Run tree problems
go run main.go stackqueue   # Run stack/queue problems
go run main.go searching    # Run searching problems
go run main.go sorting      # Run sorting algorithms
go run main.go dp           # Run dynamic programming
```

### Option 2: Run Individual Problems
```bash
# Run individual problem files
go run arrays/two_sum.go
go run strings/valid_anagram.go
go run linkedlists/reverse_linked_list.go
go run trees/max_depth_binary_tree.go
```

Note: The individual files are currently packages, so you'll need to modify them to have `package main` and a `main()` function, or create wrapper files.

### Option 3: Run Examples
```bash
go run examples/runner.go
```

### Option 4: Import as Packages (Recommended)

Create a new main file that imports the packages:

```go
package main

import (
    "golang-interview/arrays"
    "golang-interview/strings"
    // ... other imports
)

func main() {
    arrays.RunTwoSumExamples()
    strings.RunValidAnagramExamples()
    // ... call other functions
}
```

## Testing

### Run All Tests
```bash
go test ./...
```

### Run Tests for Specific Package
```bash
go test ./arrays -v
go test ./strings -v
go test ./linkedlists -v
```

### Run Tests with Coverage
```bash
go test ./arrays -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run Benchmarks
```bash
go test -bench=. ./arrays
go test -bench=BenchmarkTwoSum ./arrays
```

## Development Workflow

### 1. Adding New Problems

1. Choose a category directory (e.g., `arrays/`)
2. Create a new file (e.g., `new_problem.go`)
3. Follow the existing pattern:
   ```go
   package arrays

   /*
   Problem: [Problem Name] (LeetCode #XXX)
   [Problem description]
   */

   // Approach 1: [Description]
   // Time Complexity: O(?), Space Complexity: O(?)
   func solutionName(params...) returnType {
       // implementation
   }

   // Example usage and test cases
   func RunProblemExamples() {
       // test cases and examples
   }
   ```

4. Add tests in `[category]_test.go`
5. Update `PROBLEMS.md` with the new problem

### 2. Code Style Guidelines

- Use descriptive function names
- Include time and space complexity comments
- Provide multiple approaches when possible
- Add comprehensive test cases
- Include edge cases in examples
- Use consistent formatting

### 3. Problem Template

Copy this template for new problems:

```go
package [category]

import (
    "fmt"
    "strings"
)

/*
Problem: [Problem Name] (LeetCode #XXX)
[Detailed problem description]

Example 1:
Input: [input]
Output: [output]
Explanation: [explanation]

Constraints:
- [constraint 1]
- [constraint 2]
*/

// Approach 1: [Method name]
// Time Complexity: O(?), Space Complexity: O(?)
func problemSolution(params...) returnType {
    // Implementation
}

// Example usage and test cases
func RunProblemExamples() {
    testCases := []struct {
        input interface{}
        want  interface{}
        desc  string
    }{
        // test cases
    }
    
    fmt.Println("[Problem Name] Solutions:")
    fmt.Println(strings.Repeat("=", 40))
    
    for _, tc := range testCases {
        // Run and display results
    }
}
```

## Project Structure Explanation

```
golang-interview/
├── arrays/                 # Array and hashing problems
│   ├── two_sum.go
│   ├── group_anagrams.go
│   └── arrays_test.go
├── strings/               # String manipulation problems
│   ├── valid_anagram.go
│   └── valid_palindrome.go
├── linkedlists/           # Linked list problems
│   ├── reverse_linked_list.go
│   └── merge_two_sorted_lists.go
├── trees/                 # Binary tree problems
│   ├── max_depth_binary_tree.go
│   └── invert_binary_tree.go
├── stack-queue/           # Stack and queue problems
│   └── valid_parentheses.go
├── searching/             # Binary search problems
│   └── binary_search.go
├── sorting/               # Sorting algorithms
│   └── sorting_algorithms.go
├── dynamic-programming/   # DP problems
│   └── climbing_stairs.go
├── graphs/                # Graph algorithms (future)
├── heap/                  # Heap problems (future)
├── design/                # Data structure design (future)
├── math/                  # Mathematical problems (future)
├── utils/                 # Common utilities
│   ├── linkedlist.go
│   └── tree.go
├── examples/              # Example runners
│   └── runner.go
├── main.go                # Main entry point
├── go.mod                 # Go module file
├── README.md              # Project overview
├── PROBLEMS.md            # Detailed problem list
└── SETUP.md               # This file
```

## Tips for Learning

1. **Start Simple**: Begin with easy problems like Two Sum and Valid Anagram
2. **Understand Patterns**: Look for common patterns across similar problems
3. **Multiple Approaches**: Try to implement multiple solutions for each problem
4. **Time/Space Analysis**: Always analyze complexity before implementing
5. **Test Thoroughly**: Include edge cases and boundary conditions
6. **Practice Regularly**: Consistency is key for interview preparation

## Troubleshooting

### Common Issues

1. **Import Path Errors**: Make sure your go.mod is properly set up
2. **Package vs Main**: Remember to change `package main` when running individual files
3. **Missing Dependencies**: Run `go mod tidy` to resolve dependencies

### Getting Help

1. Check the problem comments for detailed explanations
2. Look at the test files for usage examples
3. Refer to `PROBLEMS.md` for implementation status
4. Review the NeetCode website for video explanations

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add your problem solution
4. Include tests
5. Update documentation
6. Submit a pull request

Happy coding! 🚀
