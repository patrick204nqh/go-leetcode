package main

import (
	"github.com/fatih/color"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return pathSum(root.Left, targetSum-root.Val) || pathSum(root.Right, targetSum-root.Val)
}

func main() {
	// Create a sample tree:
	//       5
	//      / \
	//     4   8
	//    /   / \
	//   11  13  4
	//  /  \      \
	// 7    2      1
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	targetSum := 22
	result := pathSum(root, targetSum)
	color.Green("Result: %v", result) // Expected output: true
}
