package main

import (
	"fmt"
)

// 上下左右のマスを探索するときに使う、差分を表す配列
var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

func main() {
	// 入力を受け取る
	var H, W int
	fmt.Scan(&H, &W)
	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	// スタートのマス、ゴールのマスを調べる
	sx, sy, gx, gy := -1, -1, -1, -1
	for x := 0; x < H; x++ {
		for y := 0; y < W; y++ {
			if string(S[x][y]) == "S" {
				sx, sy = x, y
			}
			if string(S[x][y]) == "G" {
				gx, gy = x, y
			}
		}
	}

	dist := make([][]int, H) // dist[i][j]：スタートから (i, j) にたどり着くまでの最小距離 (-1 で初期化)
	for i := range dist {
		dist[i] = make([]int, W)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	que := make([][2]int, 0) // これから調べるべき頂点を管理するキュー

	// スタート地点について、dist を更新し que に挿入する
	dist[sx][sy] = 0
	que = append(que, [2]int{sx, sy})

	// キューが空になるまで、幅優先探索を続ける
	for len(que) > 0 {
		// 次に調べるべきマスを (x, y) とおく
		x, y := que[0][0], que[0][1]
		que = que[1:]

		for i := 0; i < 4; i++ {
			// (x, y) の上下左右方向にあるマス (nx, ny) について、
			nx, ny := x+dx[i], y+dy[i]

			// そのマスが迷路の範囲内にあり、通路 かつ 未探索のマスであるならば、
			if 0 <= nx && nx < H && 0 <= ny && ny < W && string(S[nx][ny]) != "#" && dist[nx][ny] == -1 {
				// マス (nx, ny) は新たに到達可能なマスであるため、dist を更新し que に挿入する
				dist[nx][ny] = dist[x][y] + 1
				que = append(que, [2]int{nx, ny})
			}
		}
	}

	// スタートからゴールまでの最短距離を出力する (到達不可能な場合、初期値の -1 が出力される)
	fmt.Println(dist[gx][gy])
}
