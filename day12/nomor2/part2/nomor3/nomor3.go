package main

func PlayingDomino(cards [][]int, deck []int) []int {
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
}
