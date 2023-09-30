package main

import "fmt"

func main() {
	var T float64  //variable tinggi tabung
	var r float64  //variable jari-jari alas tabung
	fmt.Scanln(&T) //input tinggi tabung
	fmt.Scanln(&r) //input jari-jari alas tabung

	var luasPermukaan float64 = 2 * 3.14 * r * (r + T) //variable luas permukaan dan rumus menghitung
	fmt.Println(luasPermukaan)                         //output

}
