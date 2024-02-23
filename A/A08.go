package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	x := make([][]int, h)
	z := make([][]int, h+1)

	for i := range z {
		z[i] = make([]int, w+1)
	}

	for i := 0; i < h; i++ {
		x[i] = make([]int, w)
		for j := 0; j < w; j++ {
			var tmp int
			fmt.Scan(&tmp)
			x[i][j] = tmp
		}
	}

	var q int
	fmt.Scan(&q)

	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	d := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&a[i], &b[i], &c[i], &d[i])
	}

	//横の累積
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			z[i][j] = z[i][j-1] + x[i-1][j-1]
		}
	}

	//縦の累積
	for j := 1; j <= w; j++ {
		for i := 1; i <= h; i++ {
			z[i][j] += z[i-1][j]
		}
	}

	//回答
	for i := 0; i < q; i++ {
		fmt.Println(z[c[i]][d[i]] + z[a[i]-1][b[i]-1] - z[a[i]-1][d[i]] - z[c[i]][b[i]-1])
	}
}
