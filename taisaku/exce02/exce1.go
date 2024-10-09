//二分探索、標準関数

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := -1

	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	fmt.Println(index)

	if index < len(nums) && nums[index] == target {
		fmt.Printf("見つかりました：インデックス %d\n", index)
	} else {
		fmt.Println("見つかりませんでした")
	}
}
