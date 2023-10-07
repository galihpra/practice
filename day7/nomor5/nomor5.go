package main

import "fmt"

func RemoveDuplicates(array []int) int {
	var peta = make(map[int]bool)

	for _, i := range array {
		peta[i] = true
	}

	var array2 []int
	for j := range peta {
		array2 = append(array2, j)
	}

	return len(array2)
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9}))
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))
}
