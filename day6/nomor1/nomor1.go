package main

import "fmt"

func playWithAsterik(n int) string {
	var a string = "*"
	for i := 1; i <= n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	return a
}

func main() {
	playWithAsterik(5)
}
