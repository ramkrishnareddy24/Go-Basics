package main

import (
	"fmt"
)

func main() {
	fact:=factorial(3)
	fmt.Println(fact)
}

func factorial(number int) int {
	// res := 1

	// for i:=1; i<=number; i++ {
	// 	res = res+1 
	// }
	// return res
	if number==0 {
		return 1;
	}
	return number * factorial(number-1)
	 
}