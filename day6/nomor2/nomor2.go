package main

import (
	"fmt"
)

func isGenap(N int) string {
	var a string = ""
	if N%2 == 0 {
		fmt.Print("Z")
	} else if N%3 == 0 {
		fmt.Print("X")
	} else {
		fmt.Print("Y")
	}
	return a
}

func DrawXYZ(N int) string {
	var a string = ""
	for i := 1; i <= N*N; i++ {
		fmt.Print(isGenap(i), " ")
		if i != 1 && i%N == 0 {
			fmt.Println()
		}
	}
	return a
}
func main() {
	fmt.Println(DrawXYZ(3))
}
