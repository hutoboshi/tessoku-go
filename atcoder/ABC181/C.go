package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type Point struct {
	x int
	y int
}

func main() {
	scanner.Split(bufio.ScanWords)

	//nを入力として受け取る
	n := inputInt()

	//座標を格納するスライスを作成
	points := make([]Point, n)

	//n回のループで座標を入力して受け取り、スライスに追加する
	for i := 0; i < n; i++ {
		x := inputInt()
		y := inputInt()
		points[i] = Point{x, y}
	}

	//3点が一直線上にあるかどうかをチェックするループ
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				x1, y1 := points[i].x, points[i].y
				x2, y2 := points[j].x, points[j].y
				x3, y3 := points[k].x, points[k].y
				x1 -= x3
				x2 -= x3
				y1 -= y3
				y2 -= y3
				if x1*y2 == x2*y1 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}

func inputInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func inputString() string {
	scanner.Scan()
	text := scanner.Text()
	return text
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
