package main

import "fmt"

func main() {
	var N, M int

	_, err := fmt.Scan(&N, &M)
	if err != nil {
		fmt.Print(err)
	}

	a := make([]int, N)

	var b, c, d int

	for i := 0; i < M; i++ {
		fmt.Scan(&b, &c, &d)
		for j := b - 1; j < c; j++ {
			a[j] = d
		}
	}

	for i := 0; i < N; i++ {
		fmt.Print(a[i], " ")
	}
}
