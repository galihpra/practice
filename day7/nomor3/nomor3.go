package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	var peta = make(map[int]bool)
	for _, i := range arrayB {
		peta[i] = true
	}
	var beda []int
	// var sama []int

	for _, i := range arrayA {
		if _, ok := peta[i]; ok {
			// sama = append(sama, i)
		} else {
			beda = append(beda, i)
		}
	}

	return beda
}
func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}))
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))
}
