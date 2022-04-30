package main

import (
	"fmt"
	"math"
)

//输入：
//[4,5,1,6,2,7,3,8],4
//复制
//返回值：
//[1,2,3,4]
//复制
//说明：
//返回最小的4个数即可，返回[1,3,2,4]也可以
func main() {

	nums := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(GetLeastNumbers_Solution(nums, 4))
}

//复杂度较高
func GetLeastNumbers_Solution(input []int, k int) []int {
	if len(input) == 0 || k > len(input) {
		return []int{}
	}
	var res []int
	arr := input
	h := map[int]int{}
	for j := 0; j < k; j++ {
		min := math.MaxInt32
		for i := 0; i < len(arr); i++ { // 找到最小的
			if arr[i] < min {
				min = arr[i]
				h[arr[i]] = i
			}
		}
		res = append(res, min) // 将最小的存起来
		// 将已经选出的数从数组里删掉
		index := h[min]
		arr1 := arr[0:index]
		if index < len(arr)-1 {
			arr = arr[index+1:]
			arr = append(arr, arr1...)
		} else {
			arr = arr1
		}

	}
	return res
}
