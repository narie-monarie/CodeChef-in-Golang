package main

import "fmt"

func main() {

	var withdrawlAmt int
	fmt.Scan(&withdrawlAmt)

	var currBalance float64
	fmt.Scan(&currBalance)

	if withdrawlAmt%5 == 0 && float64(withdrawlAmt)+0.5 <= currBalance {
		currBalance = currBalance - float64(withdrawlAmt) - 0.5
		fmt.Printf("%0.2f", currBalance)

	} else {
		fmt.Printf("%0.2f", currBalance)

	}

}
