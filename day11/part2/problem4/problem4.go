package main

import "fmt"

func RomanNumerals(value int) string {
	var peta = []struct {
		Angka  int
		Simbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	for _, mapping := range peta {
		for value >= mapping.Angka {
			result += mapping.Simbol
			value -= mapping.Angka
		}
	}
	return result
}

func main() {
	fmt.Println(RomanNumerals(6))    // VI
	fmt.Println(RomanNumerals(9))    // IX
	fmt.Println(RomanNumerals(23))   // XXIII
	fmt.Println(RomanNumerals(2021)) // MMXXI
	fmt.Println(RomanNumerals(1646)) // MDCXLVI
}
