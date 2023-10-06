package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(j*i, "  ")
		}
		fmt.Println()
	}
}

func main() {
	cetakTablePerkalian(9)
}
