package main

import "fmt"

func main() {
	var A int
	var B int

	_, err := fmt.Scanln(&A, &B)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(A + B)
	}
}
