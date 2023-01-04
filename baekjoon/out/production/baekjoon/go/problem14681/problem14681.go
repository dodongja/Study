package main

import (
	"fmt"
)

func pickQuadrant(x int, y int) int {
	if x > 0 {
		if y > 0 {
			return 1
		} else {
			return 4
		}
	} else {
		if y > 0 {
			return 2
		} else {
			return 3
		}
	}
}

func main() {
	var x int
	var y int

	_, err := fmt.Scanf("%d\n %d", &x, &y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pickQuadrant(x, y))
	}

}
