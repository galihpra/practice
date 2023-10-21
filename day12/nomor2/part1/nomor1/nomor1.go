package main

func prima(number int) []int {
	var n = []int{2, 3, 5}
	for i := 6; i <= number*number; i++ {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			n = append(n, i)
		}
	}
	return n
}

func PrimeX(number int) int {
	var x = prima(number)

	return x[number-1]
}

func main() {

}
