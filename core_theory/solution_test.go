package core_theory

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic match",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Zeroes in input",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "Target equals double one element",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "No solution should return nil",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: nil,
		},
		{
			name:     "Large input array",
			nums:     generateLargeInput(1000000, 42),
			target:   1999998,
			expected: []int{999999, 999998}, // Last two elements sum to 1999998
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := two_sum(tt.nums, tt.target)

			if !reflect.DeepEqual(got, tt.expected) && !reflect.DeepEqual(reverse(got), tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func generateLargeInput(n int, target int) []int {
	arr := make([]int, n)
	for i := 0; i < n-2; i++ {
		arr[i] = i
	}
	arr[n-2] = target / 2
	arr[n-1] = target - arr[n-2]
	return arr
}

func reverse(s []int) []int {
	return []int{s[1], s[0]}
}
