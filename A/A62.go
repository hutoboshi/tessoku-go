// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var scanner = bufio.NewScanner(os.Stdin)
// var G [][]int
// var ans map[int]bool

// func main() {
// 	//入力の読み取りに使用するScannerの設定
// 	scanner.Split(bufio.ScanWords)

// 	//頂点数n,連結数mを入力から受け取る
// 	n, m := nextInt(), nextInt()

// 	//答えになる配列を作成する
// 	ans = make(map[int]bool)
// 	for i := 1; i <= n; i++ {
// 		ans[i] = false
// 	}

// 	// 隣接リストを初期化
// 	G = make([][]int, n+1)
// 	for i := range G {
// 		G[i] = make([]int, 0)
// 	}

// 	//辺の情報を入力から受け取って隣接リストを作成する
// 	for i := 0; i < m; i++ {
// 		a, b := nextInt(), nextInt()
// 		G[a] = append(G[a], b)
// 		G[b] = append(G[b], a)
// 	}

// 	//深さ優先探索を実行して結果を出力する
// 	dfs(1, -1)

// 	for i := 1; i <= n; i++ {
// 		if !ans[i] {
// 			fmt.Println("The graph is not connected.")
// 			return
// 		}
// 	}
// 	fmt.Println("The graph is connected.")
// }

// func dfs(crr, pre int) {
// 	//現在の頂点を結果に追加
// 	ans[crr] = true

// 	//隣接する書く頂点について再帰的に探索する
// 	for _, nxt := range G[crr] {
// 		if nxt != pre {
// 			dfs(nxt, crr)
// 			ans[crr] = true
// 		}
// 	}
// }

// // 次の整数を読み取るための関数
// func nextInt() int {
// 	scanner.Scan() //次のトークン（整数）を読み取る
// 	num, _ := strconv.Atoi(scanner.Text())
// 	return num
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var G [][]int
var visited []bool

func main() {
	// 入力の読み取りに使用するScannerの設定
	reader := bufio.NewReader(os.Stdin)

	// 頂点数と辺の数を入力から受け取る
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	// 辺の情報を入力から受け取って隣接リストを作成する
	G = make([][]int, N+1)
	for i := 1; i <= M; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	// 深さ優先探索を実行して結果を出力する
	visited = make([]bool, N+1)
	dfs(1)

	// 連結かどうかの判定（answer = true のとき連結）
	answer := true
	for i := 1; i <= N; i++ {
		if !visited[i] {
			answer = false
			break
		}
	}

	// 答えの出力
	if answer {
		fmt.Println("The graph is connected.")
	} else {
		fmt.Println("The graph is not connected.")
	}
}

func dfs(pos int) {
	visited[pos] = true
	for _, i := range G[pos] {
		if !visited[i] {
			dfs(i)
		}
	}
}
