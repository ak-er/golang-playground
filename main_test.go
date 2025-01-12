package main

import (
	"math/rand"
	"testing"
)

func TestFindMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive Numbers",
			input:    []int{1, 3, 5, 4, 6, 7},
			expected: 7,
		},
		{
			name:     "Negative Numers",
			input:    []int{-8, -665, -1, -87, -65},
			expected: -1,
		},
		{
			name:     "Mixed positive and negative numbers",
			input:    []int{-10, 15, -5, 30, 0},
			expected: 30,
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "All elements the same",
			input:    []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "Very large numbers",
			input:    []int{2147483647, -2147483648, 1000000000, 999999999},
			expected: 2147483647,
		},
		{
			name:     "Very small numbers (negative range)",
			input:    []int{-1000000, -999999, -2147483648, -1},
			expected: -1,
		},
		{
			name:     "Mixed duplicates with a single maximum",
			input:    []int{-1, -1, 0, 2, 2, 3, 3, 3, 10, 3, 3},
			expected: 10,
		},
		{
			name:     "Alternating positive and negative values",
			input:    []int{-1, 1, -2, 2, -3, 3, -4, 4, -5, 5},
			expected: 5,
		},
		{
			name:     "Descending sorted slice",
			input:    []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0, -10},
			expected: 100,
		},
		{
			name:     "Ascending sorted slice",
			input:    []int{-100, -50, -10, 0, 10, 20, 30, 40, 50, 60, 100},
			expected: 100,
		},
		{
			name:     "Zeros only",
			input:    []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Extremely large slice",
			input:    generateLargeSlice(1000000, 999999),
			expected: 999999,
		},
		{
			name:     "Negative edge values with zero",
			input:    []int{-2147483648, -100, -1, 0, -2147483648},
			expected: 0,
		},
		{
			name:     "Randomized unordered numbers",
			input:    []int{42, 1, 99, 88, 23, -5, 99, 0, -42},
			expected: 99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindMax(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %d but got %d", test.input, test.expected, result)
			}
		})
	}
}

func TestFindMaxEmptySlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty slice but did not panic")
		}
	}()
	FindMax([]int{})
}

func generateLargeSlice(size, max int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max)
	}
	slice[rand.Intn(size)] = max
	return slice
}
