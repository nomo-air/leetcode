package _455

import (
	"fmt"
	"sort"
	"testing"
)

/*
先尝试满足最贪心的小朋友
时间复杂度: O(nlogn)
空间复杂度: O(1)
*/
func findContentChildren(g, s ints) int {

	sort.Sort(g)
	sort.Sort(s)

	gi := len(g) - 1
	si := len(s) - 1
	res := 0
	for gi >= 0 && si >= 0 {
		if s[si] >= g[gi] {
			res++
			si--
		}
		gi--
	}

	return res
}

/*
先尝试满足最不贪心的小朋友
时间复杂度: O(nlogn)
空间复杂度: O(1)
*/

func findContentChildren1(g, s ints) int {

	sort.Sort(g)
	sort.Sort(s)

	gi := 0
	si := 0
	res := 0
	for gi < g.Len() && si < s.Len() {
		if s[si] >= g[gi] {
			res++
			gi++
		}
		si++
	}

	return res
}

type ints []int

func (s ints) Len() int {
	return len(s)
}

func (s ints) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Test455(t *testing.T) {
	fmt.Printf("先尝试满足最贪心的小朋友:%+v \n", findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Printf("先尝试满足最不贪心的小朋友:%+v \n", findContentChildren1([]int{1, 2}, []int{1, 2, 3}))
}
