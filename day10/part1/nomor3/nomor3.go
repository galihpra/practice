package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}

func primaSegiEmpat(wide, high, start int) string {
	if wide <= 0 || high <= 0 {
		return ""
	}

	result := ""
	current := start + 1
	var sum int

	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			for !isPrime(current) {
				current++
			}
			sum += current
			result += fmt.Sprintf("%d ", current)
			current++
		}
		result += "\n"
	}
	fmt.Print(result)
	fmt.Println(sum)
	return result
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)

}
