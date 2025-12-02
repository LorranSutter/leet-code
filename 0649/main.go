package main

import "fmt"

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
	fmt.Println(predictPartyVictory("R") == "Radiant")
	fmt.Println(predictPartyVictory("RRR") == "Radiant")
	fmt.Println(predictPartyVictory("D") == "Dire")
	fmt.Println(predictPartyVictory("RD") == "Radiant")
	fmt.Println(predictPartyVictory("RDD") == "Dire")
	fmt.Println(predictPartyVictory("DDRRR") == "Dire")
	fmt.Println(predictPartyVictory("DRRDRDRDRDDRDRDR") == "Radiant")
}
