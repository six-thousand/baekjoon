package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var N int
var array [][]int
var check []bool
var answer float64
var wr *bufio.Writer

func startLink(idx, count int) {
	if count == N/2 {
		sub()
		return
	}
	for i := idx; i < N; i++ {
		if !check[i] {
			check[i] = true
			startLink(i+1, count+1)
			check[i] = false
		}
	}
}

func sub() {
	teamStart := 0
	teamLink := 0

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if check[i] && check[j] {
				teamStart += array[i][j]
				teamStart += array[j][i]
			} else if !check[i] && !check[j] {
				teamLink += array[i][j]
				teamLink += array[j][i]
			}
		}
	}
	val := math.Abs(float64(teamStart) - float64(teamLink))

	if val == 0 {
		fmt.Fprint(wr, val)
		wr.Flush()
		os.Exit(0)
	}
	answer = math.Min(val, answer)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)

	defer wr.Flush()

	fmt.Fscan(rd, &N)
	array = make([][]int, N)
	check = make([]bool, N)
	answer = math.MaxFloat64

	for i := 0; i < N; i++ {
		array[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(rd, &array[i][j])
		}
	}
	startLink(0, 0)
	fmt.Fprintln(wr, int(answer))
}
