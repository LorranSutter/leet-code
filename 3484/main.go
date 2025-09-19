package main

import (
	"fmt"
	"strconv"
)

type Coords struct {
	col int
	row int
}

type Spreadsheet struct {
	Cells map[Coords]int
}

func Constructor(rows int) Spreadsheet {
	sh := Spreadsheet{
		Cells: make(map[Coords]int),
	}

	return sh
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	col := int(cell[0]) - 65
	row, _ := strconv.Atoi(cell[1:])

	this.Cells[Coords{row: row - 1, col: col}] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	col := int(cell[0]) - 65
	row, _ := strconv.Atoi(cell[1:])

	this.Cells[Coords{row: row - 1, col: col}] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	cell1, cell2 := this.GetFormulaValues(formula)

	value1 := this.Value(cell1)
	value2 := this.Value(cell2)

	return value1 + value2
}

func (this *Spreadsheet) Print() {
	for _, row := range this.Cells {
		fmt.Println(row)
	}
}

func (this *Spreadsheet) GetFormulaValues(formula string) (string, string) {
	countToPlus := 0
	for _, s := range formula {
		if s == '+' {
			break
		}
		countToPlus++
	}

	return formula[1:countToPlus], formula[countToPlus+1:]
}

func (this *Spreadsheet) Value(str string) int {
	num, err := strconv.Atoi(str)
	if err == nil { // It is just a number
		return num
	}

	col := int(str[0]) - 65
	row, _ := strconv.Atoi(str[1:])

	return this.Cells[Coords{row: row - 1, col: col}]
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */

func main() {
	sh := Constructor(15)
	// sh.GetValue("=A1+B4")
	sh.SetCell("A11", 123)
	sh.SetCell("Z3", 123)
	sh.SetCell("G7", 123)
	sh.SetCell("F4", 123)
	sh.SetCell("W14", 123)
	sh.Print()

	fmt.Println(sh.GetValue("=A11+G7"))
	fmt.Println(sh.GetValue("=6+Z3"))

	fmt.Println(sh.GetFormulaValues("=A11+G7"))
	fmt.Println(sh.GetFormulaValues("=6+Z3"))
}
