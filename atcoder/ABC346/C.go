// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	var n, k int
// 	fmt.Scan(&n, &k)

// 	arrA := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arrA[i])
// 	}

// 	sort.Ints(arrA)

// 	ans:=0
// 	var tmp int = arrA[0]
// 	if arrA[0] <=k{
// 	if arrA[0] != 1{
// 		for i:=0;i<arrA[0];i++{
// 			ans += i
// 		}
// 	}
// 	for i:=0;i<n;i++{

// 	}
// }
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var N, K int
// 	fmt.Scan(&N, &K)

// 	// 整数列Aを読み込む
// 	A := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		fmt.Scan(&A[i])
// 	}

// 	// Aに含まれない整数の総和を計算する
// 	sum := 0
// 	appeared := make(map[int]bool)
// 	for _, num := range A {
// 		appeared[num] = true
// 	}

// 	for i := 1; i <= K; i++ {
// 		if !appeared[i] {
// 			sum += i
// 		}
// 	}

// 	fmt.Println(sum)
// }

package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	// 1からKまでの整数を出現チェックするためのマップを作成
	appeared := make(map[int]bool)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		appeared[num] = true
	}

	// 1からKまでの整数のうち、Aに含まれないものの総和を計算
	sum := 0
	for i := 1; i <= K; i++ {
		if !appeared[i] {
			sum += i
		}
	}

	fmt.Println(sum)
}
