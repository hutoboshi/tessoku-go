package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// NとKを入力
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	// 友達の情報を格納するスライスを作成
	friends := make([][2]int, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		// 友達の情報を入力してスライスに追加
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		friends[i] = [2]int{a, b}
	}

	// クイックソートで友達を住んでいる村の順にソート
	sort.Slice(friends, func(i, j int) bool {
		return friends[i][0] < friends[j][0]
	})

	//現在の村を初期化し、Kを加算
	nowVillage := 0
	nowVillage += k

	// 各友達の情報を確認して村を移動
	for i := 0; i < n; i++ {
		friendVillage := friends[i][0]
		friendMoney := friends[i][1]

		// 友達の住んでいる村が現在の村以下であれば、村を移動してお金を加算
		if friendVillage <= nowVillage {
			nowVillage += friendMoney
		} else {
			//村を移動できなくなったらループを終了
		}
	}

	fmt.Println(nowVillage)

}
