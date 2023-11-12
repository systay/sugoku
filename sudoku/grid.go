package sudoku

import (
	"fmt"
	"strings"
)

type Grid [][]int

func newGrid() Grid {
	out := make([][]int, N)
	for i := range out {
		out[i] = make([]int, N)
	}
	return out
}

const N = 9

func (g Grid) get(row, col int) int {
	return g[row][col]
}
func (g Grid) set(row, col, val int) {
	g[row][col] = val
}

func Solve(g Grid, row, col int) bool {
	if col == N {
		col = 0
		row++
	}
	if row == N {
		return true
	}
	if g.get(row, col) > 0 {
		// if we already have a value, we can move on
		return Solve(g, row, col+1)
	}
	for num := 1; num <= N; num++ {
		if g.isValid(row, col, num) {
			g.set(row, col, num)
			if Solve(g, row, col+1) {
				return true
			}
		}
		g.set(row, col, 0)
	}
	return false
}

func (g Grid) isValid(row, col, value int) bool {
	for i := 0; i < N; i++ {
		if g.get(i, col) == value {
			// check the row
			return false
		}
		if g.get(row, i) == value {
			// the column
			return false
		}
	}

	startRow := row - row%3
	startCol := col - col%3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if g.get(startRow+r, startCol+c) == value {
				// the cell
				return false
			}
		}
	}

	return true
}

func (g Grid) String() string {
	builder := strings.Builder{}

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			val := g.get(r, c)
			if val == 0 {
				builder.WriteString(".")
			} else {
				builder.WriteString(fmt.Sprintf("%d", val))
			}
			if c%3 == 2 {
				builder.WriteString(" ")
			}
		}
		builder.WriteString("\n")
		if r%3 == 2 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
