package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int
var sudoku [][]int
var wr *bufio.Writer

func DFS(row, col int) {
	if col == N {
		DFS(row+1, 0)
	}

	if row == N {
		for _, s := range sudoku {
			for _, item := range s {
				fmt.Fprintf(wr, "%d ", item)
			}
			fmt.Fprintln(wr)
		}
		wr.Flush()
		os.Exit(0)
	}

	if sudoku[row][col] == 0 {
		for i := 1; i <= N; i++ {
			if possibillity(row, col, i) {
				sudoku[row][col] = i
				DFS(row, col+1)
			}
		}
		sudoku[row][col] = 0
		return
	}

	DFS(row, col+1)
}

func possibillity(row, col, value int) bool {

	// 같은 행에 있는 원소들 중 겹치는 열 원소가 있는지 검사
	for i := 0; i < N; i++ {
		if sudoku[row][i] == value {
			return false
		}
	}

	// 같은 열에 있는 원소들 중 겹치는 행 원소가 있는지 검사
	for i := 0; i < N; i++ {
		if sudoku[i][col] == value {
			return false
		}
	}

	setRow := (row / 3) * 3
	setCol := (col / 3) * 3

	for i := setRow; i < setRow+3; i++ {
		for j := setCol; j < setCol+3; j++ {
			if sudoku[i][j] == value {
				return false
			}
		}
	}

	return true
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	N = 9

	sudoku = make([][]int, N)

	for i := 0; i < N; i++ {
		sudoku[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(rd, &sudoku[i][j])
		}
	}
	DFS(0, 0)
}
