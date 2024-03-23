/*
幅優先探索
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
	n, m := nextInt(), nextInt()

	//隣接リストを作成するためのマップ
	connect := make(map[int][]int)
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		connect[a] = append(connect[a], b)
		connect[b] = append(connect[b], a)
	}

	// 各頂点のカウンタを初期化するためのマップ 初期値は0にする
	counter := make(map[int]int)
	counter[1] = 0
	for i := 2; i <= n; i++ {
		counter[i] = -1
	}

	//幅優先探索のためのキューを初期化
	que := []int{1}

	//幅優先探索
	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		for _, to := range connect[now] {
			if counter[to] == -1 {
				counter[to] = counter[now] + 1
				que = append(que, to) //キューに追加
			}
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Println(counter[i])
	}
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}
