package main

import "fmt"

func AddSubtract() {
	var a, b, c int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Scanf("%d %d", &b, &c)

		fmt.Println(10*b + 90*c)

	}

}

func main() {
	AddSubtract()

}
