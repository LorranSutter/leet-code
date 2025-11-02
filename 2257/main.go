package main

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	not_guarded := m*n - len(guards) - len(walls)
	barriers := map[[2]int]struct{}{}
	guarded := map[[2]int]struct{}{}

	for _, guard := range guards {
		key := [2]int{guard[0], guard[1]}
		barriers[key] = struct{}{}
	}

	for _, wall := range walls {
		key := [2]int{wall[0], wall[1]}
		barriers[key] = struct{}{}
	}

	for _, guard := range guards {
		// North
		key := [2]int{guard[0], guard[1]}
		for i := guard[0] - 1; i >= 0; i-- {
			key[0] = i
			if _, ok := barriers[key]; ok {
				break
			}
			if _, ok := guarded[key]; !ok {
				guarded[key] = struct{}{}
			}
		}

		// East
		key = [2]int{guard[0], guard[1]}
		for i := guard[1] + 1; i < n; i++ {
			key[1] = i
			if _, ok := barriers[key]; ok {
				break
			}
			if _, ok := guarded[key]; !ok {
				guarded[key] = struct{}{}
			}
		}

		// South
		key = [2]int{guard[0], guard[1]}
		for i := guard[0] + 1; i < m; i++ {
			key[0] = i
			if _, ok := barriers[key]; ok {
				break
			}
			if _, ok := guarded[key]; !ok {
				guarded[key] = struct{}{}
			}
		}

		// West
		key = [2]int{guard[0], guard[1]}
		for i := guard[1] - 1; i >= 0; i-- {
			key[1] = i
			if _, ok := barriers[key]; ok {
				break
			}
			if _, ok := guarded[key]; !ok {
				guarded[key] = struct{}{}
			}
		}
	}

	return not_guarded - len(guarded)
}

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}) == 7)
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}) == 4)
}
