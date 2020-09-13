package _454

import (
	"fmt"
	"testing"
)

/**
时间复杂度: O(n^2)
空间复杂度: O(n^2)
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	record := make(map[int]int, 0)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			record[C[i]+D[j]]++
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if _, ok := record[0-A[i]-B[j]]; ok {
				res += record[0-A[i]-B[j]]
			}
		}
	}

	return res
}

func Test454(t *testing.T) {
	s := fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}) // 2
	fmt.Printf("%+v", s)
}
