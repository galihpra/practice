package main

import "fmt"

func FindMinAndMax(arr []int) string {
	var idxMin int
	var idxMax int
	var min int = arr[0]
	var max int = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			idxMin = i
		}
		if arr[i] > max {
			max = arr[i]
			idxMax = i
		}
	}
	return fmt.Sprintf("min: %d, index: %d,max: %d,index: %d", min, idxMin, max, idxMax)
}
func main() {
}
