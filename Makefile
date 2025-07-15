# Golang Interview DSA Problems Makefile
# This Makefile provides convenient targets to run different problem categories and individual problems

.PHONY: help all arrays strings linkedlists trees stackqueue searching sorting dp clean test benchmark format lint

# Default target
.DEFAULT_GOAL := help

# Colors for output
GREEN := \033[0;32m
YELLOW := \033[1;33m
BLUE := \033[0;34m
RED := \033[0;31m
NC := \033[0m # No Color

## Help target - shows all available targets
help:
	@echo "$(GREEN)🚀 Golang Interview DSA Problems - Makefile$(NC)"
	@echo "$(BLUE)================================================$(NC)"
	@echo ""
	@echo "$(YELLOW)📚 Problem Categories:$(NC)"
	@echo "  make arrays          - Run all array problems"
	@echo "  make strings         - Run all string problems" 
	@echo "  make linkedlists     - Run all linked list problems"
	@echo "  make trees           - Run all tree problems"
	@echo "  make stackqueue      - Run all stack & queue problems"
	@echo "  make searching       - Run all searching problems"
	@echo "  make sorting         - Run all sorting algorithms"
	@echo "  make dp              - Run all dynamic programming problems"
	@echo "  make all             - Run all categories"
	@echo ""
	@echo "$(YELLOW)🔧 Development:$(NC)"
	@echo "  make test            - Run all tests"
	@echo "  make test-arrays     - Run tests for arrays package"
	@echo "  make test-verbose    - Run tests with verbose output"
	@echo "  make benchmark       - Run performance benchmarks"
	@echo "  make format          - Format all Go code"
	@echo "  make lint            - Run Go linter"
	@echo "  make clean           - Clean build artifacts"
	@echo ""
	@echo "$(YELLOW)📋 Individual Problems:$(NC)"
	@echo "  make two-sum         - Run Two Sum problem"
	@echo "  make group-anagrams  - Run Group Anagrams problem"
	@echo "  make contains-dup    - Run Contains Duplicate problem"
	@echo "  make product-array   - Run Product of Array Except Self"
	@echo "  make valid-anagram   - Run Valid Anagram problem"
	@echo "  make valid-palindrome - Run Valid Palindrome problem"
	@echo "  make reverse-ll      - Run Reverse Linked List problem"
	@echo "  make merge-lists     - Run Merge Two Sorted Lists problem"
	@echo "  make max-depth       - Run Maximum Depth of Binary Tree"
	@echo "  make invert-tree     - Run Invert Binary Tree problem"
	@echo "  make valid-parens    - Run Valid Parentheses problem"
	@echo "  make binary-search   - Run Binary Search problem"
	@echo "  make sorting-algos   - Run Sorting Algorithms"
	@echo "  make climbing-stairs - Run Climbing Stairs problem"
	@echo ""
	@echo "$(YELLOW)📖 Examples:$(NC)"
	@echo "  make arrays          # Run all array problems with solutions"
	@echo "  make two-sum         # Run just the Two Sum problem"
	@echo "  make test            # Run all unit tests"
	@echo "  make benchmark       # Run performance benchmarks"

# Problem Categories
arrays:
	@echo "$(GREEN)🔢 Running Array Problems...$(NC)"
	@go run main.go arrays

strings:
	@echo "$(GREEN)📝 Running String Problems...$(NC)"
	@go run main.go strings

linkedlists:
	@echo "$(GREEN)🔗 Running Linked List Problems...$(NC)"
	@go run main.go linkedlists

trees:
	@echo "$(GREEN)🌳 Running Tree Problems...$(NC)"
	@go run main.go trees

stackqueue:
	@echo "$(GREEN)📚 Running Stack & Queue Problems...$(NC)"
	@go run main.go stackqueue

searching:
	@echo "$(GREEN)🔍 Running Searching Problems...$(NC)"
	@go run main.go searching

sorting:
	@echo "$(GREEN)🔄 Running Sorting Problems...$(NC)"
	@go run main.go sorting

dp:
	@echo "$(GREEN)🧮 Running Dynamic Programming Problems...$(NC)"
	@go run main.go dp

all:
	@echo "$(GREEN)🎯 Running All Problems...$(NC)"
	@go run main.go all

# Individual Problem Targets
two-sum:
	@echo "$(BLUE)Running Two Sum Problem...$(NC)"
	@cd arrays && go run two_sum.go || go run ../main.go arrays | grep -A 20 "Two Sum Problem:"

group-anagrams:
	@echo "$(BLUE)Running Group Anagrams Problem...$(NC)"
	@cd arrays && go run group_anagrams.go || go run ../main.go arrays | grep -A 20 "Group Anagrams Problem:"

contains-dup:
	@echo "$(BLUE)Running Contains Duplicate Problem...$(NC)"
	@cd arrays && go run contains_duplicate.go || go run ../main.go arrays | grep -A 20 "Contains Duplicate Problem:"

product-array:
	@echo "$(BLUE)Running Product of Array Except Self Problem...$(NC)"
	@cd arrays && go run product_except_self.go || go run ../main.go arrays | grep -A 20 "Product of Array Except Self Problem:"

valid-anagram:
	@echo "$(BLUE)Running Valid Anagram Problem...$(NC)"
	@cd strings && go run valid_anagram.go || go run ../main.go strings | grep -A 20 "Valid Anagram Problem:"

valid-palindrome:
	@echo "$(BLUE)Running Valid Palindrome Problem...$(NC)"
	@cd strings && go run valid_palindrome.go || go run ../main.go strings | grep -A 20 "Valid Palindrome Problem:"

reverse-ll:
	@echo "$(BLUE)Running Reverse Linked List Problem...$(NC)"
	@cd linkedlists && go run reverse_linked_list.go || go run ../main.go linkedlists | grep -A 20 "Reverse Linked List Problem:"

merge-lists:
	@echo "$(BLUE)Running Merge Two Sorted Lists Problem...$(NC)"
	@cd linkedlists && go run merge_two_sorted_lists.go || go run ../main.go linkedlists | grep -A 20 "Merge Two Sorted Lists Problem:"

max-depth:
	@echo "$(BLUE)Running Maximum Depth of Binary Tree Problem...$(NC)"
	@cd trees && go run max_depth_binary_tree.go || go run ../main.go trees | grep -A 20 "Maximum Depth of Binary Tree Problem:"

invert-tree:
	@echo "$(BLUE)Running Invert Binary Tree Problem...$(NC)"
	@cd trees && go run invert_binary_tree.go || go run ../main.go trees | grep -A 20 "Invert Binary Tree Problem:"

valid-parens:
	@echo "$(BLUE)Running Valid Parentheses Problem...$(NC)"
	@cd stack-queue && go run valid_parentheses.go || go run ../main.go stackqueue | grep -A 20 "Valid Parentheses Problem:"

binary-search:
	@echo "$(BLUE)Running Binary Search Problem...$(NC)"
	@cd searching && go run binary_search.go || go run ../main.go searching | grep -A 20 "Binary Search Problem:"

sorting-algos:
	@echo "$(BLUE)Running Sorting Algorithms...$(NC)"
	@cd sorting && go run sorting_algorithms.go || go run ../main.go sorting

climbing-stairs:
	@echo "$(BLUE)Running Climbing Stairs Problem...$(NC)"
	@cd dynamic-programming && go run climbing_stairs.go || go run ../main.go dp | grep -A 20 "Climbing Stairs Problem:"

# Testing Targets
test:
	@echo "$(GREEN)🧪 Running All Tests...$(NC)"
	@go test ./... -v

test-arrays:
	@echo "$(GREEN)🧪 Running Array Tests...$(NC)"
	@go test ./arrays -v

test-strings:
	@echo "$(GREEN)🧪 Running String Tests...$(NC)"
	@go test ./strings -v

test-linkedlists:
	@echo "$(GREEN)🧪 Running Linked List Tests...$(NC)"
	@go test ./linkedlists -v

test-trees:
	@echo "$(GREEN)🧪 Running Tree Tests...$(NC)"
	@go test ./trees -v

test-stackqueue:
	@echo "$(GREEN)🧪 Running Stack & Queue Tests...$(NC)"
	@go test ./stack-queue -v

test-searching:
	@echo "$(GREEN)🧪 Running Searching Tests...$(NC)"
	@go test ./searching -v

test-sorting:
	@echo "$(GREEN)🧪 Running Sorting Tests...$(NC)"
	@go test ./sorting -v

test-dp:
	@echo "$(GREEN)🧪 Running Dynamic Programming Tests...$(NC)"
	@go test ./dynamic-programming -v

test-verbose:
	@echo "$(GREEN)🧪 Running All Tests (Verbose)...$(NC)"
	@go test ./... -v -race

test-coverage:
	@echo "$(GREEN)🧪 Running Tests with Coverage...$(NC)"
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

# Benchmark Targets
benchmark:
	@echo "$(GREEN)⚡ Running Benchmarks...$(NC)"
	@go test ./... -bench=. -benchmem

benchmark-arrays:
	@echo "$(GREEN)⚡ Running Array Benchmarks...$(NC)"
	@go test ./arrays -bench=. -benchmem

# Development Targets
format:
	@echo "$(GREEN)🎨 Formatting Go Code...$(NC)"
	@go fmt ./...

lint:
	@echo "$(GREEN)🔍 Running Go Linter...$(NC)"
	@command -v golangci-lint >/dev/null 2>&1 || { echo "$(RED)golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(NC)"; exit 1; }
	@golangci-lint run

vet:
	@echo "$(GREEN)🔍 Running Go Vet...$(NC)"
	@go vet ./...

mod-tidy:
	@echo "$(GREEN)📦 Tidying Go Modules...$(NC)"
	@go mod tidy

mod-download:
	@echo "$(GREEN)📦 Downloading Go Modules...$(NC)"
	@go mod download

# Build Targets
build:
	@echo "$(GREEN)🔨 Building Main Application...$(NC)"
	@go build -o bin/golang-interview main.go

build-all:
	@echo "$(GREEN)🔨 Building All Applications...$(NC)"
	@mkdir -p bin
	@go build -o bin/golang-interview main.go
	@go build -o bin/examples examples/runner.go

# Installation Targets
install-deps:
	@echo "$(GREEN)📦 Installing Dependencies...$(NC)"
	@go mod download
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Documentation Targets
docs:
	@echo "$(GREEN)📚 Generating Documentation...$(NC)"
	@command -v godoc >/dev/null 2>&1 || { echo "$(RED)godoc not installed. Install with: go install golang.org/x/tools/cmd/godoc@latest$(NC)"; exit 1; }
	@echo "$(GREEN)Starting documentation server at http://localhost:6060$(NC)"
	@godoc -http=:6060

# Demo Targets
demo:
	@echo "$(GREEN)🎬 Running Demo...$(NC)"
	@go run examples/runner.go

demo-interactive:
	@echo "$(GREEN)🎬 Interactive Demo - Choose a category:$(NC)"
	@echo "1) Arrays  2) Strings  3) Linked Lists  4) Trees  5) All"
	@read -p "Enter choice (1-5): " choice; \
	case $$choice in \
		1) make arrays ;; \
		2) make strings ;; \
		3) make linkedlists ;; \
		4) make trees ;; \
		5) make all ;; \
		*) echo "$(RED)Invalid choice$(NC)" ;; \
	esac

# Utility Targets
clean:
	@echo "$(GREEN)🧹 Cleaning Build Artifacts...$(NC)"
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@go clean ./...

status:
	@echo "$(GREEN)📊 Project Status:$(NC)"
	@echo "$(BLUE)Go Version:$(NC) $$(go version)"
	@echo "$(BLUE)Module:$(NC) $$(go list -m)"
	@echo "$(BLUE)Dependencies:$(NC)"
	@go list -m all | grep -v "$$(go list -m)" | head -10
	@echo "$(BLUE)Test Files:$(NC) $$(find . -name "*_test.go" | wc -l)"
	@echo "$(BLUE)Go Files:$(NC) $$(find . -name "*.go" | grep -v "_test.go" | wc -l)"

init:
	@echo "$(GREEN)🚀 Initializing Project...$(NC)"
	@go mod tidy
	@make format
	@make test
	@echo "$(GREEN)✅ Project initialized successfully!$(NC)"

# Problem-specific shortcuts
ts: two-sum
ga: group-anagrams
cd: contains-dup
pa: product-array
va: valid-anagram
vp: valid-palindrome
rl: reverse-ll
ml: merge-lists
md: max-depth
it: invert-tree
vpar: valid-parens
bs: binary-search
sa: sorting-algos
cs: climbing-stairs

# Quick aliases
a: arrays
s: strings
ll: linkedlists
t: trees
sq: stackqueue
se: searching
so: sorting
d: dp

# Show current targets
list-targets:
	@echo "$(GREEN)📋 Available Makefile Targets:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(BLUE)%-20s$(NC) %s\n", $$1, $$2}'
