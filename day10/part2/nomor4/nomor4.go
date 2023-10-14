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
func main() {
	fmt.Println(MostappearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostappearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostappearItem([]string{"football", "basketball", "tenis"}))
}
