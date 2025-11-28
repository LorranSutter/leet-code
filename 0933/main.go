package main

import "fmt"

type RecentCounter struct {
	pings   []int
	pointer int
}

func Constructor() RecentCounter {
	return RecentCounter{nil, 0}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.pings = append(rc.pings, t)
	if len(rc.pings) == 1 {
		return 1
	}

	for rc.pings[rc.pointer] < t-3000 {
		rc.pointer++
	}

	return len(rc.pings) - rc.pointer
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1) == 1)
	fmt.Println(obj.Ping(100) == 2)
	fmt.Println(obj.Ping(3001) == 3)
	fmt.Println(obj.Ping(3002) == 3)
	fmt.Println(obj.Ping(7000) == 1)
}
