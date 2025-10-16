package main

import (
	"fmt"
	"leetcode/utils"
	"strings"
)

type Row struct {
	Words           []string
	TotalWordsCount int
	TotalWordWidth  int
}

func fullJustify(words []string, maxWidth int) []string {
	var output []string

	rowLength := maxWidth
	row := Row{}
	for _, word := range words {
		lenWord := len(word)
		if (lenWord + row.TotalWordsCount) <= rowLength {
			rowLength -= lenWord
			row.Words = append(row.Words, word)
			row.TotalWordsCount++
			row.TotalWordWidth += lenWord
		} else {
			justifiedRow := justifySentence(row, maxWidth, false)
			output = append(output, justifiedRow)
			// fmt.Println(word, row, rowLength)
			row = Row{[]string{word}, 1, lenWord}
			rowLength = maxWidth - len(word)
		}
		// fmt.Println(word, row, rowLength)
	}

	output = append(output, justifySentence(row, maxWidth, true))
	// for _, row := range output {
	// 	fmt.Println(row)
	// }

	return output
}

func justifySentence(row Row, maxWidth int, lastRow bool) string {
	if row.TotalWordsCount <= 1 {
		return row.Words[0] + strings.Repeat(" ", maxWidth-row.TotalWordWidth)
	}

	if lastRow {
		lastSentence := strings.Join(row.Words, " ")
		return lastSentence + strings.Repeat(" ", maxWidth-len(lastSentence))
	}

	totalSpacing := maxWidth - row.TotalWordWidth
	minSpaceWidth := totalSpacing / (row.TotalWordsCount - 1)
	carrySpaces := totalSpacing % (row.TotalWordsCount - 1)
	// fmt.Println("Space", minSpaceWidth, carrySpaces, row.TotalWordsCount)

	output := row.Words[0]
	i := 1
	for j := carrySpaces; j > 0; j-- {
		// fmt.Println("here")
		output += strings.Repeat(" ", minSpaceWidth) + " "
		output += row.Words[i]
		i++
	}
	for ; i < row.TotalWordsCount; i++ {
		// fmt.Println("here1", i)
		output += strings.Repeat(" ", minSpaceWidth)
		output += row.Words[i]
	}

	// fmt.Println(output, len(output))
	return output
}

func main() {
	fmt.Println(
		utils.EqualSlices(
			fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16),
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		),
	)
	fmt.Println(
		utils.EqualSlices(
			fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16),
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		),
	)
	fmt.Println(
		utils.EqualSlices(
			fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20),
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		),
	)
}
