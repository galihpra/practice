package main

import "fmt"

func MaxSequence(arr []int) int {
	var temp int
	var result int
	for i := 0; i < len(arr); i++ {
		temp = arr[i]
		for j := i + 1; j < len(arr)-1; j++ {
			temp += arr[j]
			if temp > result {
				result = temp
			}
		}
	}
	return result
}
func main() {
	fmt.Println(MaxSequence([]int{-2, 1 - 3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
