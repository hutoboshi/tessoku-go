package main

import (
	"fmt"
	"sort"
)

var (
	N, K int
	R    []int
)

func main() {
	fmt.Scan(&N, &K)
	R = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&R[i])
	}

	var results [][]int
	generateSequence([]int{}, 0, &results)

	// 辞書順に並べるためにソート
	sort.Slice(results, func(i, j int) bool {
		for k := 0; k < N; k++ {
			if results[i][k] != results[j][k] {
				return results[i][k] < results[j][k]
			}
		}
		return false
	})

	// 結果を出力
	for _, seq := range results {
		for i, num := range seq {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}

func generateSequence(current []int, index int, results *[][]int) {
	if len(current) == N {
		sum := 0
		for _, num := range current {
			sum += num
		}
		if sum%K == 0 {
			*results = append(*results, append([]int(nil), current...))
		}
		return
	}

	for i := 1; i <= R[index]; i++ {
		generateSequence(append(current, i), index+1, results)
	}
}
