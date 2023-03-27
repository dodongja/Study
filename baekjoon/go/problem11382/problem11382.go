package main

import "fmt"

func main() {
	var A int
	var B int
	var C int

	_, err := fmt.Scan(&A, &B, &C)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(A + B + C)
	}
}
