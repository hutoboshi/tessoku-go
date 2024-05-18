package main

import (
	"fmt"
	"sort"
)

var arrG [][]int
var ans []int

// 深さ優先探索（ABC213 D）
func dfs(crr, pre int) {
	//現在の頂点を結果に追加
	ans = append(ans, crr)

	//隣接する各頂点について再帰的に探索する
	for _, nxt := range arrG[crr] {
		if nxt != pre {
			dfs(nxt, crr)
			//戻りの値を入れる
			ans = append(ans, crr)
		}
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	arrA := make([]int, n-1)
	arrB := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		arrA[i] = io.NextInt()
		arrB[i] = io.NextInt()
	}

	arrG = make([][]int, n+1)
	for i := range arrG {
		arrG[i] = make([]int, 0)
	}

	//辺の情報arrA,arrBから隣接リストを作成
	for i := 0; i < n-1; i++ {
		arrG[arrA[i]] = append(arrG[arrA[i]], arrB[i])
		arrG[arrB[i]] = append(arrG[arrB[i]], arrA[i])
	}

	//各頂点の隣接をソートする（小さい順に回るため）
	for i := range arrG {
		sort.Ints(arrG[i])
	}

	//深さ優先探索を実行して結果を出力する
	dfs(1, -1)

	for _, val := range ans {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
