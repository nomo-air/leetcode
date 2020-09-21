package _279

import (
	"fmt"
	"testing"
)

type Pair struct {
	Num  int
	Step int
}

func numSquares(n int) int {
	queue := make([]Pair, 0)
	queue = append(queue, Pair{Num: n, Step: 0})

	visited := make([]bool, n+1)
	visited[n] = true

	for len(queue) > 0 {
		// poll
		cur := queue[0]
		queue = queue[1:]

		num := cur.Num
		step := cur.Step

		if num == 0 {
			return step
		}

		for i := 1; num-i*i >= 0; i++ {
			if !visited[num-i*i] {
				// push
				queue = append(queue, Pair{Num: num - i*i, Step: step + 1})
				visited[num-i*i] = true
			}
		}

	}
	return 0
}

func Test279(t *testing.T) {
	fmt.Printf("%+v \n", numSquares(13))
}
