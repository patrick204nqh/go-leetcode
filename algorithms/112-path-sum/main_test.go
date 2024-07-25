package main

import (
	"github.com/fatih/color"
	"reflect"
	"testing"
)

func TestPathSum(t *testing.T) {
	testCases := []struct {
		input    []interface{}
		expected bool
	}{
		// Test Case 1: Single node tree where node's value equals target sum
		{input: []interface{}{&TreeNode{Val: 5}, 5}, expected: true},

		// Test Case 2: Multiple nodes with a path that equals the target sum
		{input: []interface{}{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 1},
					},
				},
			}, 22}, expected: true},

		// Test Case 3: Multiple nodes with no path that equals the target sum
		{input: []interface{}{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}, 5}, expected: false},

		// Test Case 4: Empty tree (root is nil)
		{input: []interface{}{nil, 0}, expected: false},
	}

	for _, tc := range testCases {
		root, _ := tc.input[0].(*TreeNode)
		sum := tc.input[1].(int)

		result := pathSum(root, sum)
		if !reflect.DeepEqual(result, tc.expected) {
			color.Red("FAIL: pathSum(%v) = %v; expected %v", tc.input, result, tc.expected)
			t.Errorf("FAIL: pathSum(%v) = %v; expected %v", tc.input, result, tc.expected)
		} else {
			color.Green("PASS: pathSum(%v) = %v", tc.input, result)
		}
	}
}
