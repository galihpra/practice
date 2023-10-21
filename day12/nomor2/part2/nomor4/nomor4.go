package main

import (
	"fmt"
	"sort"
)

func MostappearItem(items []string) string {
	var result = ""
	var map1 = make(map[string]int)
	for i := 0; i < len(items); i++ {
		if _, found := map1[items[i]]; found {
			map1[items[i]]++
		} else {
			map1[items[i]] = 1
		}
	}

	var sce = make([]int, len(map1))
	for _, value := range map1 {
		sce = append(sce, value)
	}

	sort.Ints(sce)

	for j := 0; j < len(sce); j++ {
		for index, value := range map1 {
			if value == sce[j] {
				result += fmt.Sprintln(index, "->", value)
				break
			}
		}
	}
	return result

}

func main() {
}
