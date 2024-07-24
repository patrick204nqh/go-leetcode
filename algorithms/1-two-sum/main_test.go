package main

import (
    "github.com/fatih/color"
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {
    testCases := []struct {
        nums     []int
        target   int
        expected []int
    }{
        {nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
        {nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
        {nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
    }

    for _, tc := range testCases {
        result := twoSum(tc.nums, tc.target)
        if !reflect.DeepEqual(result, tc.expected) {
            color.Red("FAIL: TwoSum(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
            t.Errorf("FAIL: TwoSum(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
        } else {
            color.Green("PASS: TwoSum(%v, %d) = %v", tc.nums, tc.target, result)
        }
    }
}
