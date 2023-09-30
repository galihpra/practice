package main

import "fmt"
func main() {
	fmt.Print("Hello bre")
	fmt.Println("Hello world")
	fmt.Print("Hai guys")

	var a int
	fmt.Scanln(&a)
	fmt.Println("Hasil input a -", a)

	var b string
	fmt.Scanln(&b)
	fmt.Println("Hasil input b- ", b)
}