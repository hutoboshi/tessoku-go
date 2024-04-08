/*
inter tools
*/
package main

import "fmt"

// permute関数は、与えられたスライスの順列を生成する
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(start int, nums []int)
	backtrack = func(start int, nums []int) {
		//バックトラック関数を再起的に呼び出して順列を生成
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			//配列の要素を交換して順列を生成
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(start+1, nums)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	backtrack(0, nums)
	return result
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	//各都市間の移動時間を格納するに2次元配列
	time := make([][]int, n)
	for i := 0; i < n; i++ {
		time[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&time[i][j])
		}
	}

	ans := 0
	perm := make([]int, n-1)
	for i := 1; i < n; i++ {
		perm[i-1] = i
	}

	//順列を生成し、それぞれの順列に対して移動時間を計算し、Kと一致するかどうか確認
	perms := permute(perm)
	for _, p := range perms {
		nowTime := time[0][p[0]]
		nowCity := p[0]

		for i := 1; i < n-1; i++ {
			toCity := p[i]
			nowTime += time[nowCity][toCity]
			nowCity = toCity
		}

		nowTime += time[nowCity][0]
		if nowTime == k {
			ans++
		}
	}

	fmt.Println(ans)
}
