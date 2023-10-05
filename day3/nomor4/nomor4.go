package main

import "fmt"

func main() {
	var hargaAwal float64
	fmt.Scanln(&hargaAwal)

	var diskon float64
	fmt.Scanln(&diskon)

	var hargaDiskon float64 = (diskon / 100) * hargaAwal
	var hargaAkhir float64 = hargaAwal - hargaDiskon

	fmt.Println("Harga yang harus dibayar adalah", hargaAkhir)
}
