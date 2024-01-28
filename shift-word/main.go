package main

import (
	"fmt"
	"strings"
)

func encodeDecode(word string) string {
	var slide = 2
	var result string

	if strings.Contains(word, "encode") {
		word = strings.ReplaceAll(word, "<encode>", "")
	} else if strings.Contains(word, "decode") {
		word = strings.ReplaceAll(word, "<decode>", "")
		slide = -2
	}

	for _, char := range word {
		var tmpResult = char + rune(slide)
		if tmpResult > 'z' {
			tmpResult -= 26
		} else if tmpResult < 'a' {
			tmpResult += 26
		}

		result += string(tmpResult)
	}

	return result
}

func main() {
	// encode
	fmt.Println(encodeDecode("<encode>abcd")) // cdef
	fmt.Println(encodeDecode("<encode>xyz"))  // zab

	// decode
	fmt.Println(encodeDecode("<decode>cdef")) // abcd
	fmt.Println(encodeDecode("<decode>zab"))  // xyz
}
