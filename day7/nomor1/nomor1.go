package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var x string

	var peta = make(map[string]bool)

	var split1 = strings.Split(a, " ")
	var split2 = strings.Split(b, " ")
	if len(split1) == 0 || len(split2) == 0 {
		fmt.Println("Kata ada yang kosong")
	}
	for i := 0; i < len(split1); i++ {
		if _, ketemu := peta[split1[i]]; !ketemu {
			peta[split1[i]] = false
		}
	}
	for i := 0; i < len(split2); i++ {
		if _, ketemu := peta[split2[i]]; !ketemu {
			peta[split2[i]] = true
		}
	}
	for j, k := range peta {
		if k {
			fmt.Print(j)
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
