package _447

import (
	"fmt"
	"testing"
)

/*
时间复杂度: O(n^2)
空间复杂度: O(n)
*/
func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		// record中存储 点i到所有其他点的距离出现的频次
		record := make(map[int]int, 0)
		for j := 0; j < len(points); j++ {
			if j != i {
				// 计算距离时不进行开根运算, 以保证精度
				dis := dis(points[i], points[j])
				if _, ok := record[dis]; ok {
					record[dis] = record[dis] + 1
				} else {
					record[dis] = 1
				}
			}

			for _, v := range record {
				res += v * (v - 1)
			}
		}
	}

	return res
}

func dis(pa, pb []int) int {
	return (pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])
}

func Test447(t *testing.T) {
	n := numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}})
	fmt.Printf("%+v", n)
}
