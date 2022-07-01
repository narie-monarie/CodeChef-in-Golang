package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&c)
	for i := 0; i < c; i++ {
		fmt.Scanf("%d %d", &a, &b)
		if a > b {
			fmt.Println(a)

		} else {
			fmt.Println(b)

		}

	}
}
