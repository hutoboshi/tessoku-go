package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 入力を読み込む
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	S := make([]string, N)
	M := 0

	for i := 0; i < N; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
		if len(S[i]) > M {
			M = len(S[i])
		}
	}

	T := make([]string, M)
	for i := range T {
		T[i] = strings.Repeat("*", N)
	}

	// 条件に基づいてTを構築する
	for i := 0; i < N; i++ {
		for j := 0; j < len(S[i]); j++ {
			T[j] = T[j][:N-i-1] + string(S[i][j]) + T[j][N-i:]
		}
	}

	// 出力
	for i := 0; i < M; i++ {
		fmt.Println(strings.TrimRight(T[i], "*"))
	}
}
