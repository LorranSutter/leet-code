package main

import "leetcode/utils"

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make(map[int][]int)

	for _, req := range prerequisites {
		adjList[req[0]] = append(adjList[req[0]], req[1])
	}

	visited := make(map[int]bool)
	finished := make(map[int]bool)

	for i := range numCourses {
		if _, ok := adjList[i]; !ok {
			continue
		}

		if hasCycle(i, adjList, visited, finished) {
			return false
		}
	}

	return true
}

func hasCycle(course int, adjList map[int][]int, visited map[int]bool, finished map[int]bool) bool {
	if finished[course] {
		return true
	}
	if visited[course] {
		return false
	}
	visited[course] = true
	finished[course] = true
	for _, v := range adjList[course] {
		if hasCycle(v, adjList, visited, finished) {
			return true
		}
	}

	finished[course] = false
	return false
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: [][]int{{1, 0}}, Got: canFinish(2, [][]int{{1, 0}}), Expected: true},
		{Input: [][]int{{1, 0}, {0, 1}}, Got: canFinish(2, [][]int{{1, 0}, {0, 1}}), Expected: false},
		{Input: [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}, Got: canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}), Expected: true},
		{Input: [][]int{{1, 0}, {1, 2}, {3, 1}, {2, 3}, {2, 4}, {4, 5}, {2, 5}}, Got: canFinish(6, [][]int{{1, 0}, {1, 2}, {3, 1}, {2, 3}, {2, 4}, {4, 5}, {2, 5}}), Expected: false},
		{Input: [][]int{{1, 0}, {1, 2}, {3, 1}, {3, 2}, {2, 4}, {4, 5}, {2, 5}}, Got: canFinish(6, [][]int{{1, 0}, {1, 2}, {3, 1}, {3, 2}, {2, 4}, {4, 5}, {2, 5}}), Expected: true},
	})
}
