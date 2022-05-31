package main

import "fmt"

func main() {
	var b int
	var a []int
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &b)
		a = append(a, b)

	}
	count := 0

	for _, i := range a {
		if i >= 10 {
			count++

		}

	}

	fmt.Println(count)

}
