package main

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
}
