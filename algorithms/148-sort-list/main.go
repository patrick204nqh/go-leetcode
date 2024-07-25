package main

import (
    "fmt"
    "github.com/fatih/color"
)

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to sort a linked list using merge sort
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    // Step 1: Split the list into two halves
    mid := getMid(head)
    left := head
    right := mid.Next
    mid.Next = nil

    // Step 2: Sort each half
    left = sortList(left)
    right = sortList(right)

    // Step 3: Merge the sorted halves
    return merge(left, right)
}

// Helper function to find the middle of the linked list
func getMid(head *ListNode) *ListNode {
    var slow, fast = head, head
    var prev *ListNode

    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    return prev
}

// Helper function to merge two sorted linked lists
func merge(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }

    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }

    return dummy.Next
}

// Helper function to print the linked list
func printList(head *ListNode) {
    current := head
    result := ""
    for current != nil {
        result += fmt.Sprintf("%d -> ", current.Val)
        current = current.Next
    }
    result += "nil"
    color.Green(result)
}

func main() {
    node1 := &ListNode{Val: 4}
    node1.Next = &ListNode{Val: 2}
    node1.Next.Next = &ListNode{Val: 1}
    node1.Next.Next.Next = &ListNode{Val: 3}

    result := sortList(node1)
    printList(result) // Expected output: 1 -> 2 -> 3 -> 4 -> nil
}
