package main

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	maze := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&maze[i])
	}

	ans := 0
	for gyou := 0; gyou < H; gyou++ {
		for retu := 0; retu < W; retu++ {
			if maze[gyou][retu] == '.' {
				ans = max(ans, explore(gyou, retu, H, W, maze))
			}
		}
	}

	fmt.Println(ans)
}

func explore(start_gyou, start_retu, H, W int, maze []string) int {
	maze_count := make([][]int, H)
	for i := 0; i < H; i++ {
		maze_count[i] = make([]int, W)
		for j := 0; j < W; j++ {
			maze_count[i][j] = -1
		}
	}
	maze_count[start_gyou][start_retu] = 0

	type pair struct{ gyou, retu int }
	que := []pair{{start_gyou, start_retu}}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(que) > 0 {
		now := que[0]
		que = que[1:]
		now_gyou, now_retu := now.gyou, now.retu
		now_count := maze_count[now_gyou][now_retu]

		for _, dir := range dirs {
			next_gyou, next_retu := now_gyou+dir[0], now_retu+dir[1]
			if next_gyou >= 0 && next_gyou < H && next_retu >= 0 && next_retu < W && maze[next_gyou][next_retu] == '.' && maze_count[next_gyou][next_retu] == -1 {
				maze_count[next_gyou][next_retu] = now_count + 1
				que = append(que, pair{next_gyou, next_retu})
			}
		}
	}

	max_count := 0
	for _, row := range maze_count {
		for _, val := range row {
			max_count = max(max_count, val)
		}
	}

	return max_count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
