package __1

import (
	"fmt"
	"math/rand"
	"testing"
)

// 递归使用快速排序,对arr[l...r]的范围进行排序
func sort(arr []int, l, r int) {
	if l < 0 || r < 0 || r <= l {
		return
	}
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	swap(arr, l, rand.Intn(r-l+1)+l)

	v := arr[l]

	lt := l     // arr[l+1...lt] < v
	gt := r + 1 // arr[gt...r] > v
	i := l + 1  // arr[lt+1...i) == v
	for i < gt {
		if arr[i] < v {
			swap(arr, i, lt+1)
			i++
			lt++
		} else if arr[i] > v {
			swap(arr, i, gt-1)
			gt--
		} else { // arr[i] == v
			i++
		}
	}

	// 把v放回切入点
	swap(arr, l, lt)

	sort(arr, l, lt-1)
	sort(arr, gt, r)
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

/*
三路快速排序算法也是一个O(nlogn)复杂度的算法
可以在1秒之内轻松处理100万数量级的数据
*/
func Test_1(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort(arr, 0, len(arr)-1)
	fmt.Printf("%+v", arr)
}
