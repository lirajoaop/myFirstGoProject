package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	a, b := swap(10, 20)
	fmt.Println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
