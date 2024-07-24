package main

import (
    "github.com/fatih/color"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

// 
// Detects if a cycle exists in a linked list
// 
func linkedListCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow, fast := head, head

    // Traverse the list with two pointers
    // Slow pointer moves one step at a time
    // Fast pointer moves two steps at a time
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        // If slow and fast pointers meet, there is a cycle
        if slow == fast {
            return true
        }
    }

    // If we reach here, there is no cycle
    return false
}

func main() {
    // Create a linked list with a cycle
    node1 := &ListNode{Val: 3}
    node2 := &ListNode{Val: 2}
    node3 := &ListNode{Val: 0}
    node4 := &ListNode{Val: -4}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node2 // Creates a cycle

    result := linkedListCycle(node1)
    color.Green("Result: %v", result) // Expected output: true
}
