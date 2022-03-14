package main

import "fmt"

func nwd(x int, y int) int {
	var a, b int = x, y
	if x < y {
		a = y
		b = x
	}
	if a % b != 0 {
		return nwd(b, a % b)
	}
	return b
}

func nwd2(x int, y int) int {
	if x < y { return nwd2(y, x) }
	if x % y != 0 { return nwd2(y, x % y) }
	return y
}

func main() {
	x := 24
	y := 160

	fmt.Println(nwd(x, y))
	fmt.Println(nwd2(x, y))
}
