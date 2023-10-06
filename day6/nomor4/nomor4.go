package main

import "fmt"

func ubahHuruf(sentence string) string {
	var a string = ""
	for i := 0; i < len(sentence); i++ {
		if int(sentence[i]) > 80 {
			a += string(sentence[i] + 10 - 26)
		} else if int(sentence[i]) == 32 {
			a += " "
		} else {
			a += string(sentence[i] + 10)
		}
	}
	return a
}
func main() {
	fmt.Println(ubahHuruf("SEPULSA OKE"))
	fmt.Println(ubahHuruf("ALTERRA ACADEMY"))
	fmt.Println(ubahHuruf("INDONESIA"))
	fmt.Println(ubahHuruf("GOLANG"))
	fmt.Println(ubahHuruf("PROGRAMMER"))
}
