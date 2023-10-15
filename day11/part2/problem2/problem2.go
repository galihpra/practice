package main

import "fmt"

func BottomUp(n int) int {
	var a int = 0
	var b int = 1
	var c int = 0

	if n < 2 {
		return n
	}

	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {
	fmt.Println(BottomUp(0))  //0
	fmt.Println(BottomUp(1))  //1
	fmt.Println(BottomUp(2))  //1
	fmt.Println(BottomUp(3))  //2
	fmt.Println(BottomUp(5))  //5
	fmt.Println(BottomUp(6))  //8
	fmt.Println(BottomUp(7))  //13
	fmt.Println(BottomUp(9))  //34
	fmt.Println(BottomUp(10)) //55
}
