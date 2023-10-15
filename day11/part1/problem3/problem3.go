package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) string {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	result := 0
	i, j := 0, 0

	for i < len(dragonHead) && j < len(knightHeight) {
		if dragonHead[i] <= knightHeight[j] {
			result += knightHeight[j]
			i++
		}
		j++
	}

	if i == len(dragonHead) {
		return fmt.Sprintf("%d", result)
	}

	return "Knight Fall"
}

func main() {
	fmt.Println(DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}))    //11
	fmt.Println(DragonOfLoowater([]int{5, 10}, []int{5}))         // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})) // knight fall
	fmt.Println(DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})) // 10
}
