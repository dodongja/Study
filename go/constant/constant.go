package main

import "fmt"

func main() {
	const (
		Red   int = iota
		Blue  int = iota
		Green int = iota
	)

	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)

	const (
		C1 uint = iota + 1
		C2
		C3
	)

	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)

	const PI = 3.14

	var a int = PI * 100
	var b float32 = PI * 100

	fmt.Println(a)
	fmt.Println(b)
}
