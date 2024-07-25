package main

import (
	"github.com/fatih/color"
	"reflect"
	"testing"
)

func TestFindMinimumInRolatedSortedArray(t *testing.T) {
	testCases := []struct {
		input    []interface{}
		expected int
	}{
		// Test Case 1: Standard rotated array
		{input: []interface{}{[]int{3, 4, 5, 1, 2}}, expected: 1},

		// Test Case 2: Array rotated at the middle
		{input: []interface{}{[]int{4, 5, 6, 7, 0, 1, 2}}, expected: 0},

		// Test Case 3: Array rotated at the end
		{input: []interface{}{[]int{1, 2, 3, 4, 5, 6, 0}}, expected: 0},

		// Test Case 4: Array not rotated
		{input: []interface{}{[]int{1, 2, 3, 4, 5}}, expected: 1},

		// Test Case 5: Single element array
		{input: []interface{}{[]int{1}}, expected: 1},

		// Test Case 6: Array with two elements rotated
		{input: []interface{}{[]int{2, 1}}, expected: 1},
	}

	for _, tc := range testCases {
		nums := tc.input[0].([]int)

		result := findMinimumInRolatedSortedArray(nums)
		if !reflect.DeepEqual(result, tc.expected) {
			color.Red("FAIL: findMinimumInRolatedSortedArray(%v) = %v; expected %v", tc.input, result, tc.expected)
			t.Errorf("FAIL: findMinimumInRolatedSortedArray(%v) = %v; expected %v", tc.input, result, tc.expected)
		} else {
			color.Green("PASS: findMinimumInRolatedSortedArray(%v) = %v", tc.input, result)
		}
	}
}
