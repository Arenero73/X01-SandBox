package main

import "fmt"

func main() {
	const count = 25
	a, b := 0, 1
	for i := 0; i < count; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
