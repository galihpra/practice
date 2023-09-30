package main

import "fmt"

func main() {
	var nama string                                                                 //variable nama
	fmt.Scanln(&nama)                                                               //input nama
	fmt.Println("Hello " + nama + "! Saya golang bahasa yang sangat menyenangkan.") //output
}
