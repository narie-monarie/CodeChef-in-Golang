package main

import "fmt"

func main() {
	var a, b, y int
	fmt.Scanln(&y)
	for i := 0; i < y; i++ {
		fmt.Scanf("%d  %d", &a, &b)

		if (a+b)%2 == 0 {
			fmt.Println("YES")

		} else {
			fmt.Println("NO")

		}

	}

}
