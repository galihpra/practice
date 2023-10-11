package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var x string

	var peta = make(map[string]bool)

	var split1 = strings.Split(a, "")
	var split2 = strings.Split(b, "")
	for _, kata := range split1 {
		peta[kata] = true
	}

	for _, kata := range split2 {
		if _, ok := peta[kata]; ok {
			fmt.Print(kata)
		}
	}
	return x
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
