package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var bytes = []byte(input)

	for i := range bytes {
		if offset > 26 {
			if bytes[i]+byte(offset%26) > 122 {
				bytes[i] = bytes[i] + byte(offset%26) - 26
			} else {
				bytes[i] = bytes[i] + byte(offset%26)

			}
		} else {
			if bytes[i]+byte(offset) > 122 {
				bytes[i] = bytes[i] + byte(offset) - 26
			} else {
				bytes[i] = bytes[i] + byte(offset)
			}

		}
	}
	return string(bytes)

}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
