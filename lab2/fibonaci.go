package main

import "fmt"

func fib(n int) {
	a := 0
	b := 1
	placeholder := 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		placeholder = b
		b = a + b
		a = placeholder
	}
}

func main() {
	n := 10

	fib(n)
}
