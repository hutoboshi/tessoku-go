package main

import (
	"fmt"
	"sort"
	"strconv"
)

func g1(n int) int {
	digits := []rune(strconv.Itoa(n)) // 数字を文字列に変換し、文字列をルーン（文字）のスライスに変換
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j] // 文字列を降順にソート
	})
	sortedStr := string(digits)          // ソートされた文字列を再び文字列に変換
	result, _ := strconv.Atoi(sortedStr) // 文字列を整数に変換
	return result
}

func g2(n int) int {
	digits := []rune(strconv.Itoa(n)) // 数字を文字列に変換し、文字列をルーン（文字）のスライスに変換
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j] // 文字列を昇順にソート
	})
	sortedStr := string(digits)          // ソートされた文字列を再び文字列に変換
	result, _ := strconv.Atoi(sortedStr) // 文字列を整数に変換
	return result
}

func f(n int) int {
	return g1(n) - g2(n) // g1(n)からg2(n)を引いて結果を返す
}

func main() {
	var N, K int
	fmt.Scan(&N, &K) // 入力を受け取る
	for i := 0; i < K; i++ {
		N = f(N) // 指定された回数だけ処理を繰り返す
	}
	fmt.Println(N) // 結果を出力する
}
