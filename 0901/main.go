package main

import "fmt"

// 152 ms / 11.12 MB
// The idea here is to use a list to keep track of the stock prices.
// For each new price, we scan the list backwards until we find a price greater than the current price.
// The span is then the difference between the current index and the index of the last higher price.
// This approach is less efficient because in the worst case we may have to scan the entire list for each new price.
type StockSpanner1 struct {
	stack []int
}

func Constructor1() StockSpanner1 {
	return StockSpanner1{nil}
}

func (s1 *StockSpanner1) Next1(price int) int {
	if s1.stack == nil {
		s1.stack = []int{price}
		return 1
	}

	s1.stack = append(s1.stack, price)
	i := len(s1.stack) - 2
	for ; i >= 0; i-- {
		if s1.stack[i] > price {
			break
		}
	}

	return len(s1.stack) - i - 1
}

// 28 ms / 9.9 MB
// The idea here is to use a stack to keep track of the stock prices and their spans.
// For each new price, we pop prices from the stack until we find a price greater than the current price.
// The span is then the sum of the spans of the popped prices plus one for the current price.
// This approach is more efficient because each price is pushed and popped from the stack at most once.
type StockSpanner2 struct {
	stack [][2]int
}

func Constructor2() StockSpanner2 {
	return StockSpanner2{nil}
}

func (s2 *StockSpanner2) Next2(price int) int {
	span := 1
	for len(s2.stack) > 0 && s2.stack[len(s2.stack)-1][0] <= price {
		span += s2.stack[len(s2.stack)-1][1]
		s2.stack = s2.stack[:len(s2.stack)-1]
	}
	s2.stack = append(s2.stack, [2]int{price, span})

	return span
}

func main() {
	fmt.Println("Solution 01")
	obj1 := Constructor1()
	fmt.Println(obj1.Next1(100) == 1)
	fmt.Println(obj1.Next1(80) == 1)
	fmt.Println(obj1.Next1(60) == 1)
	fmt.Println(obj1.Next1(70) == 2)
	fmt.Println(obj1.Next1(60) == 1)
	fmt.Println(obj1.Next1(75) == 4)
	fmt.Println(obj1.Next1(85) == 6)
	fmt.Println(obj1.Next1(100) == 8)
	fmt.Println(obj1.Next1(110) == 9)
	fmt.Println(obj1.Next1(80) == 1)

	fmt.Println()

	fmt.Println("Solution 02")
	obj2 := Constructor2()
	fmt.Println(obj2.Next2(100) == 1)
	fmt.Println(obj2.Next2(80) == 1)
	fmt.Println(obj2.Next2(60) == 1)
	fmt.Println(obj2.Next2(70) == 2)
	fmt.Println(obj2.Next2(60) == 1)
	fmt.Println(obj2.Next2(75) == 4)
	fmt.Println(obj2.Next2(85) == 6)
	fmt.Println(obj2.Next2(100) == 8)
	fmt.Println(obj2.Next2(110) == 9)
	fmt.Println(obj2.Next2(80) == 1)
}
