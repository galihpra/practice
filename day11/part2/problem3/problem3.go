package main

import "fmt"

func Frog(jumps []int) int {
	var result int

	for i := 0; i < len(jumps); i++ {
		if i+2 < len(jumps) {
			if result < jumps[i+2]-jumps[i] {
				result = jumps[i+2] - jumps[i]
			}
		}
	}

	return result

}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
