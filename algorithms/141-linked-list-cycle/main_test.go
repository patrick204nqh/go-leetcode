package main

import (
    "github.com/fatih/color"
    "reflect"
    "testing"
)

// Helper function to create a linked list from a slice of values and a cycle position
func createLinkedList(values []int, pos int) *ListNode {
    if len(values) == 0 {
        return nil
    }

    nodes := make([]*ListNode, len(values))
    for i, val := range values {
        nodes[i] = &ListNode{Val: val}
        if i > 0 {
            nodes[i-1].Next = nodes[i]
        }
    }

    if pos >= 0 {
        nodes[len(nodes)-1].Next = nodes[pos]
    }

    return nodes[0]
}

func TestLinkedListCycle(t *testing.T) {
    testCases := []struct {
        input struct {
            values []int
            pos    int
        }
        expected bool
    }{
        // Cycle at node with value 2
        {input: struct{ values []int; pos int }{values: []int{3, 2, 0, -4}, pos: 1}, expected: true},
        // No cycle
        {input: struct{ values []int; pos int }{values: []int{1, 2}, pos: -1}, expected: false},
        // Single node, no cycle
        {input: struct{ values []int; pos int }{values: []int{1}, pos: -1}, expected: false},
        // Empty list, no cycle
        {input: struct{ values []int; pos int }{values: []int{}, pos: -1}, expected: false},
         // Cycle at node with value 1
        {input: struct{ values []int; pos int }{values: []int{1, 2, 3, 4, 5}, pos: 0}, expected: true},
    }

    for _, tc := range testCases {
        head := createLinkedList(tc.input.values, tc.input.pos)
        result := linkedListCycle(head)
        if !reflect.DeepEqual(result, tc.expected) {
            color.Red("FAIL: linkedListCycle(%v, %d) = %v; expected %v", tc.input.values, tc.input.pos, result, tc.expected)
            t.Errorf("FAIL: linkedListCycle(%v, %d) = %v; expected %v", tc.input.values, tc.input.pos, result, tc.expected)
        } else {
            color.Green("PASS: linkedListCycle(%v, %d) = %v", tc.input.values, tc.input.pos, result)
        }
    }
}
