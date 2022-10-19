package main

import "fmt"

func printNo(n int) {
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}
