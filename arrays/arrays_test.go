package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (Hash Map)", func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" (Brute Force)", func(t *testing.T) {
			got := twoSumBruteForce(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Has duplicate",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "No duplicates",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Multiple duplicates",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" (Optimized)", func(t *testing.T) {
			if got := containsDuplicateOptimized(tt.nums); got != tt.want {
				t.Errorf("containsDuplicateOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "With zero",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "Two elements",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name+" (With Arrays)", func(t *testing.T) {
			if got := productExceptSelfWithArrays(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelfWithArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.Run("HashMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			twoSum(nums, target)
		}
	})

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			twoSumBruteForce(nums, target)
		}
	})
}
