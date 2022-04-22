package main

import "container/heap"

func main() {

	s := []int{3, 2, 1, 5, 6, 4}

	findKthLargest(s, 2)
}
func findKthLargest(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	h := iheap{}
	heap.Init(&h)
	// 先把前k个元素构建小顶堆
	for i := 0; i < k; i++ {
		heap.Push(&h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		// 看一下堆顶元素，如果小于等于当前nums[i]就没必要替换了，减少调整次数
		top := h.Peek()
		if nums[i] > top {
			heap.Pop(&h) // 弹出堆顶
			heap.Push(&h, nums[i])
		}
	}
	// 此时堆顶元素就是第K大
	return h.Peek()
}

// 以上 也可以每次先Push，然后在判断长度是否大于K，大于K时再Pop，这样是利用堆的自调整特性保证长度为K
// 不过调整次数比增加

type iheap []int

// 实现 Sort 接口的Len, Less, Swap方法
func (h iheap) Len() int           { return len(h) }
func (h iheap) Less(i, j int) bool { return h[i] < h[j] }
func (h iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 看一下堆顶元素，不调整
func (h iheap) Peek() int {
	return h[0]
}

// 实现 container/Heap 接口的Pop,Push方法，receiver需要是指针
func (h *iheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 弹出数组最后一个元素，同时数组长度--，堆内部通过替换堆顶元素来进行 heapifyDown 调整
func (h *iheap) Pop() interface{} {
	n := (*h).Len()
	x := (*h)[n-1]  // 数组最后一个元素
	*h = (*h)[:n-1] // size--
	return x
}

// 作者：beyondyyh
// 链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/215-shu-zu-zhong-de-di-kge-zui-da-yuan-s-g6n2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
