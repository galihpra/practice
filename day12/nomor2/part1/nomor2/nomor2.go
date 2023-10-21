package main

func Fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	}
	return Fibonacci(number-2) + Fibonacci(number-1)
}
func main() {

}
