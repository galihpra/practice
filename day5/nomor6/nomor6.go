package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	//Base case
	if num == 2 || num == 3 || num == 5 {
		return true
	}
	// Check divisibility of a number
	for i := num / 2; i > 1; i-- {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func FullPrima(N int) bool {
	var result bool = false
	if isPrime(N) {
		result = true
		var numb int = N
		var digit int = 0
		for numb > 0 && result {
			// Get last digit
			digit = numb % 10
			if digit == 2 || digit == 3 ||
				digit == 5 || digit == 7 {
				numb = numb / 10
			} else {
				// When digit is not prime
				result = false
			}
		}
	}
	if result {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(FullPrima(2))
	fmt.Println(FullPrima(3))
	fmt.Println(FullPrima(11))
	fmt.Println(FullPrima(13))
	fmt.Println(FullPrima(23))
	fmt.Println(FullPrima(29))
	fmt.Println(FullPrima(37))
	fmt.Println(FullPrima(41))
	fmt.Println(FullPrima(43))
	fmt.Println(FullPrima(53))
}
