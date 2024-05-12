package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 6, 7}
	fmt.Println(binarySearchFirst(arr, 7))
}

// 二分探索(xの値が配列のどこに存在するか？)
func binarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		m := (left + right) / 2
		if x < arr[m] {
			right = m - 1
		} else if x == arr[m] {
			return m
		} else {
			left = m + 1
		}
	}
	return -1
}

// xで指定した値未満の値のindexを取得(非推奨)
func binarySearchFirst(arr []int, x int) int {
	left := 0
	right := len(arr) - 1
	result := -1 // x 以下の最初の要素のインデックスを保持する変数

	for left <= right {
		m := (left + right) / 2

		if x <= arr[m] {
			result = m
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return result
}
