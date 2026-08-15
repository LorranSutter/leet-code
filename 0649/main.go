package main

import "leetcode/utils"

func predictPartyVictory(senate string) string {
	if len(senate) == 1 {
		if senate[0] == 'R' {
			return "Radiant"
		} else {
			return "Dire"
		}
	}

	queue := []rune(senate)
	countR, countD := 0, 0
	for i := range queue {
		if queue[i] == 'R' {
			countR++
		} else {
			countD++
		}
	}

	if countR == 0 {
		return "Dire"
	}
	if countD == 0 {
		return "Radiant"
	}

	toBan := 'R'
	for len(queue) > 1 {
		toBan = 'R'
		if queue[0] == 'R' {
			toBan = 'D'
		}

		for i := 1; i < len(queue); i++ {
			if queue[i] == toBan {
				queue = append(queue[:i], queue[i+1:]...)
				queue = append(queue, queue[0])
				queue = queue[1:]
				if toBan == 'R' {
					countR--
					if countR == 0 {
						return "Dire"
					}
				} else {
					countD--
					if countD == 0 {
						return "Radiant"
					}
				}
				break
			}
		}
	}

	if queue[0] == 'R' {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func main() {
	utils.RunTests([]utils.TestCase[string]{
		{Input: "R", Got: predictPartyVictory("R"), Expected: "Radiant"},
		{Input: "RRR", Got: predictPartyVictory("RRR"), Expected: "Radiant"},
		{Input: "D", Got: predictPartyVictory("D"), Expected: "Dire"},
		{Input: "RD", Got: predictPartyVictory("RD"), Expected: "Radiant"},
		{Input: "RDD", Got: predictPartyVictory("RDD"), Expected: "Dire"},
		{Input: "DDRRR", Got: predictPartyVictory("DDRRR"), Expected: "Dire"},
		{Input: "DRRDRDRDRDDRDRDR", Got: predictPartyVictory("DRRDRDRDRDDRDRDR"), Expected: "Radiant"},
	})
}
