package main

import (
	"testing"
)

func Test_TwoSum(t *testing.T) {
	t.Run("test two sum", func(t *testing.T) {
		testCase := []struct {
			name     string
			nums     []int
			target   int
			expected []int
		}{
			{
				name:     "test is true",
				nums:     []int{1, 2, 3, 4, 5},
				target:   3,
				expected: []int{0, 1},
			},
			{
				name:     "test is true",
				nums:     []int{1, 2, 3},
				target:   5,
				expected: []int{1, 2},
			},
		}

		for _, tt := range testCase {
			t.Run(tt.name, func(t *testing.T) {
				actual := twoSum(tt.nums, tt.target)

				if len(actual) != len(tt.expected) {
					t.Errorf("expected %v but got %v", len(tt.expected), len(actual))
				}

				for i := range tt.expected {
					if tt.nums[i] != actual[i] {
						t.Errorf("expected %v bot got %v", tt.nums[i], actual[i])
					}
				}

			})
		}
	})
}
