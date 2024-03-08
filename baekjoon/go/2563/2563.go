package main

import "fmt"

func main() {
	var paper [100][100]int
	var num int
	var x int
	var y int
	var add int
	var area int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&x, &y)
		for j := 0; j < 10; j++ {
			add = y
			for k := 0; k < 10; k++ {
				paper[x][add] = 1
				add++
			}
			x++
		}
	}

	for l := 0; l < 100; l++ {
		for m := 0; m < 100; m++ {
			if paper[l][m] == 1 {
				area++
			}
		}
	}

	fmt.Print(area)
}
