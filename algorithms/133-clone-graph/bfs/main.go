package main

import (
	"github.com/fatih/color"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph clones a graph given a reference node using BFS.
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	queue := []*Node{node}
	visited[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, neighbor := range n.Neighbors {
			if _, found := visited[neighbor]; !found {
				visited[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
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
	// Example usage
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	result := cloneGraph(node1)
	printGraph(result)
}
