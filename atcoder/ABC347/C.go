// package main

// import "fmt"

// func main() {
// 	var n, a, b int
// 	fmt.Scan(&n, &a, &b)
// 	arrD := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arrD[i])
// 	}

// 	var ans bool
// 	for i := 0; i < n; i++ {
// 		ans = true
// 		for today := 0; today < a+b; today++ {
// 			tmp := (arrD[i] + today) % (a + b)
// 			if tmp == 0 {
// 				ans = false
// 			}
// 			if tmp > a {
// 				ans = false
// 			}
// 			if ans {
// 				fmt.Println("Yes")
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println("No")
// }

// package main

// import "fmt"

// func main() {
// 	var n, a, b int
// 	fmt.Scan(&n, &a, &b)
// 	arrD := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arrD[i])
// 	}

// 	var ans bool
// 	for i := 0; i < n; i++ {
// 		ans = true
// 		for today := 1; today <= a+b; today++ {
// 			tmp := (arrD[i] + today) % (a + b)
// 			if tmp == 0 || tmp > a {
// 				ans = false
// 			}
// 		}
// 		if ans {
// 			fmt.Println("Yes")
// 			return
// 		}
// 	}
// 	fmt.Println("No")
// }

// package main

// import "fmt"

// func main() {
// 	var n, a, b int
// 	fmt.Scan(&n, &a, &b)
// 	arrD := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&arrD[i])
// 	}

// 	week := a + b
// 	for i := 0; i < n; i++ {
// 		tmp := arrD[i] % week
// 		if tmp == 0 {
// 			fmt.Println("No")
// 			return
// 		}
// 		if tmp > a {
// 			fmt.Println("No")
// 			return
// 		}
// 	}
// 	fmt.Println("Yes")
// }

package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	arrD := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrD[i])
	}

	for i := 0; i < n; i++ {
		workdays := a
		holidays := b
		if arrD[i] <= workdays {
			workdays -= arrD[i]
		} else if arrD[i] <= workdays+holidays {
			holidays -= arrD[i] - workdays
			workdays = 0
		} else {
			holidays = 0
			workdays = 0
		}

		if workdays == 0 && holidays > 0 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
