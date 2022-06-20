package main

import "fmt"

func Burgers() {
	var n, a, b, count int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		if a > 1 && b > 1 {
			count++
		}

		if a > b {
			fmt.Println(b)
		} else {
			fmt.Println(a)
		}
	}
}

func main() {
	Burgers()
}
