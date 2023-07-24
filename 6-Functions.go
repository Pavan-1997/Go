package main

import "fmt"

func countdown(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

func sum(num1 int, num2 int) int {
	t := num1 + num2
	return t
}
func main() {
	countdown(10)
	fmt.Println("Sum is", sum(1, 2))
}
