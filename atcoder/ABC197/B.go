package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w, x, y int
	fmt.Scan(&h, &w, &x, &y)

	arrS := make([]string, h+2)
	arrS[0] = strings.Repeat("#", w+2)
	arrS[h+1] = strings.Repeat("#", w+2)
	for i := 1; i <= h; i++ {
		var tmp string
		fmt.Scan(&tmp)
		arrS[i] = "#" + tmp + "#"
	}

	counter := 1

	//上
	i := 1
	for {
		if arrS[x-i][y] == '#' {
			break
		}
		counter++
		i++
	}
	//下
	i = 1
	for {
		if arrS[x+i][y] == '#' {
			break
		}
		counter++
		i++
	}
	//左
	i = 1
	for {
		if arrS[x][y-i] == '#' {
			break
		}
		counter++
		i++
	}
	//右
	i = 1
	for {
		if arrS[x][y+i] == '#' {
			break
		}
		counter++
		i++
	}

	fmt.Println(counter)

}
