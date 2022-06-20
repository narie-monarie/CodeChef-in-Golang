package main

import "fmt"

func waiTime() {
	var input, a, b int
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		fmt.Scanln(&a, &b)
		if a*7 > b {
			fmt.Println(a*7 - b)
		}
	}
}

func main() {
	waiTime()

}
