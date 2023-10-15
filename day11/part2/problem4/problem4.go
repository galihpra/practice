package main

import "fmt"

func RomanNumerals(value int) string {
	var num = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result []int
	var pos int
	for value > 0 {
		for value >= num[pos] {
			value -= num[pos]
			result = append(result, num[pos])
		}
		pos++
	}

	var str string
	for i := 0; i < len(result); i++ {
		if result[i] == 1000 {
			str += "M"
		} else if result[i] == 900 {
			str += "CM"
		} else if result[i] == 500 {
			str += "D"
		} else if result[i] == 400 {
			str += "CD"
		} else if result[i] == 100 {
			str += "C"
		} else if result[i] == 90 {
			str += "XC"
		} else if result[i] == 50 {
			str += "L"
		} else if result[i] == 40 {
			str += "XL"
		} else if result[i] == 10 {
			str += "X"
		} else if result[i] == 9 {
			str += "IX"
		} else if result[i] == 5 {
			str += "V"
		} else if result[i] == 4 {
			str += "IV"
		} else if result[i] == 1 {
			str += "I"
		}
	}

	return str
}

func main() {
	fmt.Println(RomanNumerals(6))    //VI
	fmt.Println(RomanNumerals(9))    //IX
	fmt.Println(RomanNumerals(23))   //XXIII
	fmt.Println(RomanNumerals(2021)) //MMXXI
	fmt.Println(RomanNumerals(1646)) //MDCXLVI
}
