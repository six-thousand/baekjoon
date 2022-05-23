package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var N, min, max int
var seqs []int
var operators []int

func DFS(num, idx int) {
	if idx == N {
		max = int(math.Max(float64(max), float64(num)))
		min = int(math.Min(float64(min), float64(num)))
		return
	}
	for i := range operators {
		if operators[i] > 0 {
			operators[i]--
			switch i {
			case 0:
				DFS(num+seqs[idx], idx+1)
				break
			case 1:
				DFS(num-seqs[idx], idx+1)
				break
			case 2:
				DFS(num*seqs[idx], idx+1)
				break
			case 3:
				DFS(num/seqs[idx], idx+1)
				break
			}
			operators[i]++
		}
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	max = math.MinInt64
	min = math.MaxInt64

	fmt.Fscan(rd, &N)
	seqs = make([]int, N)
	operators = make([]int, 4)

	for i := range seqs {
		fmt.Fscan(rd, &seqs[i])
	}

	for i := range operators {
		fmt.Fscan(rd, &operators[i])
	}
	DFS(seqs[0], 1)

	fmt.Fprintf(wr, "%d\n%d", max, min)
}
