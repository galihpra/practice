package main

import (
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) int {
	sort.Slice(productPrice, func(i, j int) bool { return productPrice[i] < productPrice[j] })
	var hasil = []int{}
	for i := 0; i < len(productPrice); i++ {
		if money >= productPrice[i] {
			money -= productPrice[i]
			hasil = append(hasil, productPrice[i])
		} else {
			break
		}
	}
	return (len(hasil))
}
func main() {
}
