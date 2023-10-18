package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) interface{} {
	// var result string = "no solution"
	for x := 0; x < a; x++ {
		for y := 0; y < a; y++ {
			for z := 0; z < a; z++ {
				if x+y+z == a && x*y*z == b {
					// result = fmt.Sprint([]int{x, y, z})
					return []int{x, y, z}
				}
			}
		}
	}
	return fmt.Sprintf("no solution")

}

func main() {
	fmt.Println(SimpleEquations(1, 2, 3))  //no solution
	fmt.Println(SimpleEquations(6, 6, 14)) // [1,2,3]
}
