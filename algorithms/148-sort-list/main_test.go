package main

import (
    "github.com/fatih/color"
    "testing"
)

// Helper function to convert a slice to a linked list
func sliceToList(nums []int) *ListNode {
    if len(nums) == 0 {
        return nil
    }
    head := &ListNode{Val: nums[0]}
    current := head
    for _, num := range nums[1:] {
        current.Next = &ListNode{Val: num}
        current = current.Next
    }
    return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
    var result []int
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}

// Helper function to compare two linked lists
func isEqual(l1, l2 *ListNode) bool {
    for l1 != nil && l2 != nil {
        if l1.Val != l2.Val {
            return false
        }
        l1 = l1.Next
        l2 = l2.Next
    }
    return l1 == nil && l2 == nil
}

func TestSortList(t *testing.T) {
    testCases := []struct {
        input    []int
        expected []int
    }{
        {input: []int{}, expected: []int{}},
        {input: []int{1}, expected: []int{1}},
        {input: []int{3, 1, 2}, expected: []int{1, 2, 3}},
        {input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
        {input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
        {input: []int{1, 3, 2, 3}, expected: []int{1, 2, 3, 3}},
    }

    for _, tc := range testCases {
        inputList := sliceToList(tc.input)
        expectedList := sliceToList(tc.expected)
        result := sortList(inputList)
        resultSlice := listToSlice(result)
        expectedSlice := listToSlice(expectedList)
        if !isEqual(result, expectedList) {
            color.Red("FAIL: sortList(%v) = %v; expected %v", tc.input, resultSlice, expectedSlice)
            t.Errorf("FAIL: sortList(%v) = %v; expected %v", tc.input, resultSlice, expectedSlice)
        } else {
            color.Green("PASS: sortList(%v) = %v", tc.input, resultSlice)
        }
    }
}
