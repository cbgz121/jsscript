package main

import (
	"fmt"
)

// 冒泡排序函数
func bubble_sort(arr []int) {
	n := len(arr)

	// 遍历所有数组元素
	for i := 0; i < n; i++ {
		// 最后 i 个元素已经就位，不需要再比较
		for j := 0; j < n-i-1; j++ {
			// 如果元素比相邻的元素大，则交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 示例使用
func main() {
	// 要排序的列表
	my_list := []int{64, 34, 25, 12, 22, 11, 90}

	// 调用冒泡排序函数
	bubble_sort(my_list)

	// 输出排序后的列表
	fmt.Println("排序后的列表:", my_list)
}
