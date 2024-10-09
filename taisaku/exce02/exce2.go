// 二分探索1から書く
package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		fmt.Println(left, right)
		mid := left + (right-left)/2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 15

	index := binarySearch(nums, target)
	if index != -1 {
		fmt.Printf("見つかりました: インデックス %d\n", index)
	} else {
		fmt.Println("見つかりませんでした")
	}
}
