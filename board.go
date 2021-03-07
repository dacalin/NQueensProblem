package main

import (
	"fmt"
)

type Board struct {
	size int
	queens_count int
	solution map[int]int
	cols map[int]int
	diag45 map[int]int
	diag135 map[int]int
}

func (board *Board) Init( size int)  {

	board.size = size
	board.queens_count = 0
	board.solution = make(map[int]int)
	board.cols = make(map[int]int)
	board.diag45 = make(map[int]int)
	board.diag135 = make(map[int]int)

	for pos:=0; pos < size; pos++ {
		board.cols[pos] = pos
	}

	for pos:= 0; pos < (size+size-1); pos++ {
		board.diag45[pos] = pos
	}

	for pos:= -size+1; pos < (size); pos++ {
		board.diag135[pos] = pos
	}
}

func (board *Board) SetQueen( row int, col int)  {
	board.queens_count += 1
	board.solution[row] = col

	delete(board.cols, col)
	delete(board.diag45, col+row)
	delete(board.diag135, col-row)
}

func (board *Board) RemoveQueen(row int, col int)  {
	delete(board.solution, row)

	board.cols[col] = col
	board.diag45[col+row] = col+row
	board.diag135[col-row] = col-row
	board.queens_count -= 1
}

func (board Board) IsValidPos( row int, col int) bool {

	_, okayC := board.cols[col]
	_, okay45 := board.diag45[col+row]
	_, okay135 := board.diag135[col-row]

	return okay45 && okay135 && okayC
}

func (board Board) Print() {
	fmt.Println()

	for r:= 0; r < board.size; r++ {

		fmt.Print(" ")

		for c := 0; c < board.size; c++ {
			v, ok := board.solution[r]
			if ok && v== c {
				fmt.Print(4, "  ")
			} else {
				fmt.Print(0, "  ")
			}
		}

		fmt.Println()
	}
}
