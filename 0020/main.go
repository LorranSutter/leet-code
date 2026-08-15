package main

import (
	"leetcode/utils"
)

// The idea here is to use a stack to keep track of the opening parentheses.
// If we encounter a closing parenthesis, we check if the last opening parenthesis is the corresponding type.
// If they don't match, we return false.
// At the end, if the stack is empty, it means all parentheses were matched correctly.
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || pairs[char] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	utils.RunTests([]utils.TestCase[bool]{
		{Input: "[", Got: isValid("["), Expected: false},
		{Input: "()", Got: isValid("()"), Expected: true},
		{Input: "()[]{}", Got: isValid("()[]{}"), Expected: true},
		{Input: "(]", Got: isValid("(]"), Expected: false},
		{Input: "([])", Got: isValid("([])"), Expected: true},
		{Input: "([)])", Got: isValid("([)])"), Expected: false},
		{Input: "([{}])", Got: isValid("([{}])"), Expected: true},
		{Input: "[({})]", Got: isValid("[({})]"), Expected: true},
	})
}
