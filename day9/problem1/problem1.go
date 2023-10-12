package main

import "fmt"

func Swap(a, b *int) (int, int) {
	temp := *a
	*a = *b
	*b = temp

	return *a, *b
}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}
