package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var G [][]int
var ans []int

func main() {
	//入力の読み取りに使用するScannerの設定
	scanner.Split(bufio.ScanWords)

	// 頂点数nを入力から受け取る
	n := nextInt()

	//隣接リストを初期化
	G = make([][]int, n+1)
	for i := range G {
		G[i] = make([]int, 0)
	}

	//辺の情報を入力から受け取って隣接リストを作成する
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	//各頂点の隣接リストをソートする
	for i := range G {
		sort.Ints(G[i])
	}

	//深さ優先探索を実行して結果を出力する
	dfs(1, -1)

	//結果を出力
	for _, val := range ans {
		fmt.Print(val, " ")
	}
	fmt.Println()

}

func dfs(crr, pre int) {
	//現在の頂点を結果に追加
	ans = append(ans, crr)

	//隣接する書く頂点について再起的に探索する
	for _, nxt := range G[crr] {
		if nxt != pre {
			dfs(nxt, crr)
			ans = append(ans, crr)
		}
	}
}

// 次の整数を読み取るための関数
func nextInt() int {
	scanner.Scan()                         //次のトークン（整数）を読み取る
	num, _ := strconv.Atoi(scanner.Text()) //文字列を例数に変換
	return num                             //整数を返す
}
