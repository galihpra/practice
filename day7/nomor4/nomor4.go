package main

import "fmt"

func findMaxSumSubArray(k int, arr []int) int {
	var jumlah int = 0

	for i := 0; i <= len(arr); i++ {
		for j := 0; j < i; j++ {
			if len(arr[j:i]) == k {
				var sub int = 0
				for _, angka := range arr[j:i] {
					sub += angka
				}
				if sub > jumlah {
					jumlah = sub
				}
			}
		}

	}
	return jumlah
}
func main() {
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))

}
