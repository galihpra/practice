package main

import (
	"fmt"
	"sort"
)

func MostappearItem(items []string) string {
	var a string
	var peta = make(map[string]int)
	for _, value := range items {
		peta[value]++
	}

	slc := make([][2]string, 0, len(peta))

	for value, count := range peta {
		slc = append(slc, [2]string{value, fmt.Sprintf("%d", count)})
	}

	sort.Slice(slc, func(i, j int) bool {
		return slc[i][1] < slc[j][1]
	})

	for _, value := range slc {
		fmt.Print(value[0], "->", value[1], " ")
	}
	return a

}
func SeringMuncul(item []string) string {
	var result = ""
	var map1 = make(map[string]int)
	for i := 0; i < len(item); i++ {
		if _, found := map1[item[i]]; found {
			map1[item[i]]++
		} else {
			map1[item[i]] = 1
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
	fmt.Println(SeringMuncul([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(SeringMuncul([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(SeringMuncul([]string{"football", "basketball", "tenis"}))
}
