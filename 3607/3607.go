package main

import (
	"fmt"
	"leetcode/utils"
)

type GraphNode struct {
	Id          int
	Operational bool
	Edges       []*GraphNode
}

// TLE
func processQueries1(c int, connections [][]int, queries [][]int) []int {
	graph := make([]*GraphNode, c+1)

	for i := range c + 1 {
		newNode := &GraphNode{Id: i, Operational: true}
		graph[i] = newNode
	}

	// Create graph connections
	for _, connection := range connections {
		graph[connection[0]].Edges = append(graph[connection[0]].Edges, graph[connection[1]])
		graph[connection[1]].Edges = append(graph[connection[1]].Edges, graph[connection[0]])
	}

	result := []int{}
	for _, query := range queries {
		if query[0] == 2 {
			graph[query[1]].Operational = false
			continue
		}

		if graph[query[1]].Operational {
			result = append(result, query[1])
			continue
		}

		visited := make(map[int]bool, c+1)
		smallest := c + 1
		DFS(graph, query[1], visited, c, &smallest)

		if smallest < c+1 {
			result = append(result, smallest)
		} else {
			result = append(result, -1)
		}
	}

	return result
}

func DFS(graph []*GraphNode, startNode int, visited map[int]bool, c int, smallest *int) {
	visited[startNode] = true
	if graph[startNode].Operational && startNode < *smallest {
		*smallest = startNode
	}

	for _, e := range graph[startNode].Edges {
		if !visited[e.Id] {
			DFS(graph, e.Id, visited, c, smallest)
		}
	}
}

// Disjoint set union
type DSUNode struct {
	size           int
	representative int
}

type DSU struct {
	dsu map[int]DSUNode
}

// func (dsi )

func (dsu DSU) find(u int) int {
	if u == dsu.dsu[u].representative {
		return u
	} else {
		dsu.dsu[u].representative = dsu.find(dsu.dsu[u].representative)
		return dsu.dsu[u].representative
	}
}

func processQueries2(c int, connections [][]int, queries [][]int) []int {
	operational := make(map[int]bool)
	dsu := make(map[int]DSUNode)

	for i := range c {
		operational[i+1] = true
	}

	// Create connections dsu
	for _, connection := range connections {
		in_set := false
		for _, connection_set := range connection_sets {
			if _, ok := connection_set[connection[0]]; ok {
				connection_set[connection[1]] = struct{}{}
				in_set = true
				break
			}
			if _, ok := connection_set[connection[1]]; ok {
				connection_set[connection[0]] = struct{}{}
				in_set = true
				break
			}
		}

		if !in_set {
			new_connection_set := map[int]struct{}{
				connection[0]: {},
				connection[1]: {},
			}
			connection_sets = append(connection_sets, new_connection_set)
		}
	}

	// result := []int{}
	// for _, query := range queries {
	// 	if query[0] == 2 {
	// 		operational[query[1]] = false
	// 		continue
	// 	}

	// 	if operational[query[1]] {
	// 		result = append(result, query[1])
	// 		continue
	// 	}

	// 	smallest := c + 1
	// 	for _, connection_set := range connection_sets {
	// 		if _, ok := connection_set[query[1]]; ok {
	// 			// Look for minimum
	// 			for station := range connection_set {
	// 				if operational[station] && station < smallest {
	// 					smallest = station
	// 				}
	// 			}

	// 			break
	// 		}
	// 	}

	// 	if smallest == c+1 {
	// 		result = append(result, -1)
	// 	} else {
	// 		result = append(result, smallest)
	// 	}
	// }

	return result
}

func main() {
	fmt.Println("Solution 01")
	connections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	result := processQueries1(5, connections, queries)
	fmt.Println(utils.DeepEqualSlices(result, []int{3, 2, 3}))

	connections = [][]int{}
	queries = [][]int{{1, 1}, {2, 1}, {1, 1}}
	result = processQueries1(3, connections, queries)
	fmt.Println(utils.DeepEqualSlices(result, []int{1, -1}))

	fmt.Println()

	fmt.Println("Solution 02")
	connections = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries = [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	result = processQueries2(5, connections, queries)
	fmt.Println(utils.DeepEqualSlices(result, []int{3, 2, 3}))

	connections = [][]int{}
	queries = [][]int{{1, 1}, {2, 1}, {1, 1}}
	result = processQueries2(3, connections, queries)
	fmt.Println(utils.DeepEqualSlices(result, []int{1, -1}))

	connections = [][]int{{1, 2}, {2, 3}, {4, 5}, {5, 6}, {3, 4}}
	queries = [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}, {1, 6}}
	result = processQueries2(6, connections, queries)
	fmt.Println(utils.DeepEqualSlices(result, []int{3, 2, 3, 6}))
}
