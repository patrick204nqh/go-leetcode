package main

import (
	"github.com/fatih/color"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph clones a graph given a reference node using DFS.
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

// clone performs the DPS and clones the nodes.
func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// If the node has been visited, return its clone.
	if clonedNode, found := visited[node]; found {
		return clonedNode
	}

	// Clone the node and add it to the visited map.
	clonedNode := &Node{Val: node.Val}
	visited[node] = clonedNode

	// Recursively clone all the neighbors.
	for _, neighbor := range node.Neighbors {
		clonedNode.Neighbors = append(clonedNode.Neighbors, clone(neighbor, visited))
	}

	return clonedNode
}

func printGraph(node *Node) {
	visited := make(map[*Node]bool)
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil || visited[n] {
			return
		}
		visited[n] = true
		color.Green("Node %d: {", n.Val)
		for _, neighbor := range n.Neighbors {
			color.Green("  Neighbor %d", neighbor.Val)
		}
		color.Green("}")
		for _, neighbor := range n.Neighbors {
			dfs(neighbor)
		}
	}
	dfs(node)
}

func main() {
	// Create a graph: 1 - 2 - 3 - 4 - 1
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	result := cloneGraph(node1)

	color.Green("Cloned Graph:")
	printGraph(result)
}
