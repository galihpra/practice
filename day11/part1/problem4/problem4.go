package main

import "fmt"

func BinarySearch(array []int, x int) int {
	var res int
	for i := 0; i < len(array); i++ {
		if array[i] == x {
			res = i
		}
	}
	if res == 0 {
		res = -1
	}
	return res
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))                  //2
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))                 //3
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))  //6
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)) //-1
}
