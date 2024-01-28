package main

import "fmt"

func findDuplicationNumber(numbers []int) {
	var show = make(map[int]int)
	var result []int

	for _, num := range numbers {
		show[num]++
		if show[num] == 2 {
			result = append(result, num)
		}
	}

	fmt.Println(result)
}

func main() {
	findDuplicationNumber([]int{1, 1})
	findDuplicationNumber([]int{1, 2, 3, 4})
	findDuplicationNumber([]int{1, 2, 1, 2, 2, 3, 4, 5, 5, 5, 5})
}
