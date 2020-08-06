package _202

import (
	"fmt"
	"testing"
)

func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func Test202(t *testing.T) {
	b := isHappy(1)
	fmt.Printf("%+v", b)
}
