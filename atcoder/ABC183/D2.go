/*
幅優先探索
時間内に処理が終わる
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords) //スペース区切りで入力を読み取る
	n, q := nextInt(), nextInt()   //超点数とクエリ数を取得

	//隣接リストを作成するためのマップ
	connect := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		connect[a] = append(connect[a], b)
		connect[b] = append(connect[b], a)
	}

	//各頂点のカウンタを初期化するためのマップ
	counter := make(map[int]int)
	for i := 0; i < q; i++ {
		p, x := nextInt(), nextInt()
		counter[p] += x
	}

	//幅優先探索のためのキューを初期化
	que := []int{1}	
	//訪問済みの頂点を管理するためのマップ
	visited := make(map[int]bool)
	visited[1] = true
	//幅優先探索
	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		nowNumber := counter[now]
		for _, to := range connect[now] {
			if !visited[to] {
				counter[to] += nowNumber //カウンタに加算
				visited[to] = true       //訪問済みに設定
				que = append(que, to)    //キューに追加
			}
		}
	}

	//結果を出力
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(counter[i])
	}
	fmt.Println()
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}
