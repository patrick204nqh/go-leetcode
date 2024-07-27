package main

import (
	"reflect"
	"testing"

	"github.com/fatih/color"
)

// Helper function to format the graph structure
func formatGraph(node *Node) string {
	if node == nil {
		return color.New(color.FgYellow).Sprint("nil")
	}
	visited := make(map[*Node]bool)
	var dfs func(*Node) string
	dfs = func(n *Node) string {
		if n == nil || visited[n] {
			return ""
		}
		visited[n] = true
		result := color.New(color.FgCyan).Sprintf("Node %d: {", n.Val)
		for _, neighbor := range n.Neighbors {
			result += color.New(color.FgCyan).Sprintf(" Neighbor %d", neighbor.Val)
		}
		result += " }\n"
		for _, neighbor := range n.Neighbors {
			result += dfs(neighbor)
		}
		return result
	}
	return dfs(node)
}

func TestCloneGraph(t *testing.T) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	testCases := []struct {
		input    *Node
		expected *Node
	}{
		// Test case 1: Empty graph
		{input: nil, expected: nil},
		// Test case 2: Single node graph
		{input: &Node{Val: 1}, expected: &Node{Val: 1}},
		// Test case 3: Simple graph
		{input: node1, expected: node1},
	}

	for i, tc := range testCases {
		result := cloneGraph(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			color.Red("FAIL: Test case %d", i+1)
			color.Red("Input:\n%s", formatGraph(tc.input))
			color.Red("Expected:\n%s", formatGraph(tc.expected))
			color.Red("Got:\n%s", formatGraph(result))
		} else {
			color.Green("PASS: Test case %d", i+1)
			color.Green("Input:\n%s", formatGraph(tc.input))
			color.Green("Got:\n%s", formatGraph(result))
		}
	}
}
