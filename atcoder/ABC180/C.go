package main

import (
	"fmt"
	"sort"
)

// 与えられた数の約数のリストを作成してソートして返す関数
func divisorList(n int) []int {
	var divisor []int

	i := 1
	for i*i <= n {
		if n%i == 0 {
			divisor = append(divisor, i) //約数を追加
			if i != n/i {                //平方根でない場合、逆数も増加
				divisor = append(divisor, n/i)
			}
		}
		i++
	}
	sort.Ints(divisor) //約数をソート
	return divisor
}

func main() {
	var n int
	fmt.Scan(&n) //入力を受け取る

	ans := divisorList(n) //約数のリストを取得

	for _, x := range ans { //約数のリストを出力
		fmt.Println(x)
	}
}
