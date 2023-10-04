package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input nama siswa lebih dari 1 kata
	inputReader := bufio.NewReader(os.Stdin)
	studentName, _ := inputReader.ReadString('\n')

	// input nilai siswa
	var studentScore int
	fmt.Scanln(&studentScore)

	// menampilkan output nama siswa
	fmt.Print("Nama siswa: ", studentName)

	// if statement untuk menentukan kategori nilai siswa
	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("Nilai A")
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("Nilai B+")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("Nilai B")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("Nilai C")
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Format nilai tidak sesuai") // jika nilai yang diinputkan selain bilangan 0-100
	}
}
