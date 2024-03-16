// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x float64

// 	fmt.Scan(&x)

// 	fmt.Println(math.Ceil(x / 10))
// }

package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	// x が負の場合は符号を反転させてから計算する
	if x < 0 {
		x = -x
		fmt.Println(-(x / 10))
	} else {
		fmt.Println((x + 9) / 10)
	}
}
