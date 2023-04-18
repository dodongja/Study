package main

import (
	"fmt"
)

func main() {
	var N, M int

	_, err := fmt.Scan(&N, &M)
	if err != nil {
		fmt.Print(err)
	}

	a := make([]int, N)

	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	var b, c int

	for i := 0; i < M; i++ {
		fmt.Scan(&b, &c)
		var s []int
		s = a[b-1 : c]

		var revs []int

		for k := range s {
			revs = append(revs, s[len(s)-1-k])
		}

		for j := b - 1; j < c; j++ {
			a[j] = revs[0]
			revs = revs[1:]
		}
	}

	for i := range a {
		fmt.Print(a[i], " ")
	}
}
