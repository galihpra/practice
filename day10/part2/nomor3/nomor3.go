package main

import "fmt"

func playingDomino(cards [][]int, deck []int) []int {
	var result []int
	for i := 0; i < len(cards); i++ {
		// for j := 0; j < len(deck); j++ {
		if cards[i][0] == deck[0] || cards[i][1] == deck[0] {
			result = append(result, cards[i][0])
			result = append(result, cards[i][1])
		} else if cards[i][0] == deck[1] || cards[i][1] == deck[1] {
			result = append(result, cards[i][0])
			result = append(result, cards[i][1])
		}
		// }
	}
	var finalResult []int
	if len(result) > 0 {
		finalResult = append(finalResult, result[0], result[1])
	}
	return finalResult
}
func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
