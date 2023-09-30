package main

import "fmt"

func main() {
	var alas float64   //variabel alas
	var tinggi float64 ////variabel tinggi

	fmt.Scanln(&alas)                      //input alas
	fmt.Scanln(&tinggi)                    //input tinggi
	var luas float64 = (alas * tinggi) / 2 //variabel luas dan rumus menghitung

	fmt.Println(luas) //output
}
