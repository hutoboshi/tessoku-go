/*
幅優先探索
処理自体は正解だが
TLEになる
D2.goにもう一つの解を書く
*/

package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	//つながっている頂点のリストを作る
	connect := make([][]int, n+1)

	//受け取り
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		//aからb、bからaがつながっている
		connect[a] = append(connect[a], b)
		connect[b] = append(connect[b], a)
	}

	//各頂点のカウンタを用意
	counter := make([]int, n+1)

	//受け取り
	for i := 0; i < q; i++ {
		var p, x int
		fmt.Scan(&p, &x)
		//pのcounterにxを足す
		counter[p] += x
	}

	//筒を用意する
	que := make([]int, 0)

	//スタート地点=1を入れる
	que = append(que, 1)

	//訪問済みチェックを作る
	//Falseなら訪問、Trueなら訪問済み
	visited := make([]bool, n+1)
	//1=スタート地点は訪問済みにする
	visited[1] = true

	//筒が空になるまで
	for len(que) > 0 {
		//筒の左から取り出してnowに格納
		now := que[0]
		que = que[1:]
		//nowカウンタの値をnowNumberに格納
		nowNumber := counter[now]
		//nowからつながっている頂点を順に回る=to
		for _, to := range connect[now] {
			//もし訪問済みでなければ
			if !visited[to] {
				//conunterにnowのカウンタの値を足す
				counter[to] += nowNumber
				//訪問済みチェックつける
				visited[to] = true
				//筒に入れる
				que = append(que, to)
			}
		}
	}

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(counter[1])
		} else {
			fmt.Print(" ")
			fmt.Print(counter[i])
		}
	}
	fmt.Println()
}
