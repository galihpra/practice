package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) interface{} {
	if a+b == c {
		return "no solution"
	} else {
		return []int{1, 2, 3}
	}

}

func main() {
	fmt.Println(SimpleEquations(1, 2, 3))  //no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // [1,2,3]
}
