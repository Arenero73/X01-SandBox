package main

import "fmt"

// Fibonacci returns an iterator that yields the first n Fibonacci numbers.
func Fibonacci(n int) func(func(uint64) bool) {
	return func(yield func(uint64) bool) {
		var a, b uint64 = 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	for v := range Fibonacci(20) {
		fmt.Println(v)
	}
}
