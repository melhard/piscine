package piscine

import (
	"github.com/01-edu/z01"
)

type Board [8][8]bool

func (b *Board) PlaceQueen(x, y int) {
	b[x][y] = true
}

func (b *Board) RemoveQueen(x, y int) {
	b[x][y] = false
}

func (b *Board) IsValid(x, y int) bool {
	for i := 0; i < 8; i++ {
		if b[i][y] {
			return false
		}
	}

	for j := 0; j < 8; j++ {
		if b[x][j] {
			return false
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b[i][j] {
			return false
		}
	}

	for i, j := x+1, y+1; i < 8 && j < 8; i, j = i+1, j+1 {
		if b[i][j] {
			return false
		}
	}

	for i, j := x-1, y+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if b[i][j] {
			return false
		}
	}

	for i, j := x+1, y-1; i < 8 && j >= 0; i, j = i+1, j-1 {
		if b[i][j] {
			return false
		}
	}

	return true
}

func SolveQueens(b *Board, x int, solutions *[92]Board) int {
	var solutionCount int

	var solve func(x int)
	solve = func(x int) {
		if x == 8 {
			var copy Board
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					copy[i][j] = b[i][j]
				}
			}
			solutions[solutionCount] = copy
			solutionCount++
			return
		}

		for y := 0; y < 8; y++ {
			if b.IsValid(x, y) {
				b.PlaceQueen(x, y)
				solve(x + 1)
				b.RemoveQueen(x, y)
			}
		}
	}

	solve(0)
	return solutionCount
}

func EightQueens() {
	var solutions [92]Board
	var b Board

	numSolutions := SolveQueens(&b, 0, &solutions)

	for i := 0; i < numSolutions; i++ {
		for j := 0; j < 8; j++ {
			for k := 0; k < 8; k++ {
				if solutions[i][j][k] {
					z01.PrintRune(rune(k + '1'))
				}
			}
		}
		z01.PrintRune('\n')
	}
}
