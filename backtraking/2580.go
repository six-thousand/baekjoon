package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int
var array [][]int
var wr *bufio.Writer

func sudoku(row, col int) {
	// 행을 다 채우면 다음 열 시작
	if col == N {
		sudoku(row+1, 0)
	}
	if row == N {
		for i := range array {
			for j := range array[i] {
				fmt.Fprintf(wr, "%d ", array[i][j])
			}
			fmt.Fprintln(wr)
		}
		wr.Flush()
		os.Exit(0)
		return
	}
	if array[row][col] == 0 {
		for i := 1; i < N; i++ {
			if possibillity(row, col, i) {
				array[row][col] = 1
				sudoku(row, col+1)
			}
		}
		array[row][col] = 0
		return
	}
	sudoku(row, col+1)
}

func possibillity(row, col, value int) bool {

	// 같은 열에 똑같은 값이 있는지 확인
	for i := 0; i < N; i++ {
		if array[row][i] == value {
			return false
		}
	}

	// 같은 행에 똑같은 값이 있는지 확인
	for i := 0; i < N; i++ {
		if array[i][col] == value {
			return false
		}
	}

	// 같은 구역에 똑같은 값이 있는지 확인
	startSectionRow := row / 3 * 3
	startSectionCol := col / 3 * 3

	for i := startSectionRow; i < startSectionRow+3; i++ {
		for j := startSectionCol; j < startSectionCol+3; j++ {
			if array[i][j] == value {
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
	defer wr.Flush()

	array = make([][]int, N)

	for i := 0; i < N; i++ {
		array[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(rd, &array[i][j])
		}
	}
	sudoku(0, 0)

}
