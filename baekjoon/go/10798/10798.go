package main

import (
	"fmt"
)

func main() {
	var a [5][15]string

	for i := 0; i < 5; i++ {
		var input string
		fmt.Scanln(&input)
		for j := 0; j < len(input); j++ {
			a[i][j] = string(input[j])
		}
	}

	for i := 0; i < 15; i++ {
		for j := 0; j < 5; j++ {
			if len(a[j][i]) != 0 {
				fmt.Print(a[j][i])
			}
		}
	}

}
