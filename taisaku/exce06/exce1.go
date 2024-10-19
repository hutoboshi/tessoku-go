package main

import (
	"fmt"
	"sort"
)

func main() {
	// 切手の額面のリスト
	stamps := []int{1, 2, 3, 3, 7, 9, 10}

	//切手の組み合わせによって表せる金額のセット
	sums := make(map[int]bool)

	//組み合わせの合計を求める
	sums[0] = true

	for _, stamp := range stamps {
		tmp := make(map[int]bool)
		for sum := range sums {
			tmp[sum+stamp] = true
		}
		for k := range tmp {
			sums[k] = true
		}
	}

	arrS := make([]int, 0)
	for tmp2, _ := range sums {
		arrS = append(arrS, tmp2)
	}

	sort.Ints(arrS)
	fmt.Println(arrS)
	fmt.Println(len(sums))

}
