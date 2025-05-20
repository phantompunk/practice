package kata

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	// Define the test cases.  Each test case has a name, an input array,
	// and the expected output.
	testCases := []struct {
		name   string
		input  []int
		expect [][]int
	}{
		{
			name:   "empty array",
			input:  []int{},
			expect: [][]int{},
		},
		{
			name:   "less than 3 elements",
			input:  []int{1, 2},
			expect: [][]int{},
		},
		{
			name:   "no triplets",
			input:  []int{1, 2, 3},
			expect: [][]int{},
		},
		{
			name:   "one triplet",
			input:  []int{-1, 0, 1, 2, -1, -4},
			expect: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:   "all zeros",
			input:  []int{0, 0, 0, 0},
			expect: [][]int{{0, 0, 0}},
		},
		{
			name:   "duplicate triplets",
			input:  []int{-2, 0, 0, 2, 2},
			expect: [][]int{{-2, 0, 2}},
		},
		// {
		// 	name:   "negative numbers",
		// 	input:  []int{-1, -1, -1, 0, 1, 2, -2, -1},
		// 	expect: [][]int{{-2, -1, 3}, {-1, -1, 2}, {-1, 0, 1}},
		// },
		// {
		// 	name:  "large input",
		// 	input: []int{-4, -2, -2, -2, 0, 0, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
		// 	expect: [][]int{
		// 		{-4, -2, 6},
		// 		{-4, 0, 4},
		// 		{-4, 1, 3},
		// 		{-4, 2, 2},
		// 		{-2, -2, 4},
		// 		{-2, 0, 2},
		// 	},
		// },
	}

	// Iterate over the test cases and run each one.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested.
			actual := threeSum(tc.input)

			// Use DeepEqual to compare the actual output to the expected output.
			// This is necessary because the order of the triplets matters in
			// this problem, and DeepEqual handles nested slices correctly.
			if !reflect.DeepEqual(actual, tc.expect) {
				t.Errorf("threeSum(%v) = %v, expected %v", tc.input, actual, tc.expect)
			}
		})
	}
}
