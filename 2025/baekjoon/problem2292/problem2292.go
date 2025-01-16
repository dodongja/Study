package main

import "fmt"

func main() {
	// 1
	//  2*1 + 2*2 6
	//  2*3 + 3*2 12
	//  2*5 + 4*2 18
	//  2*7 + 5*2 24

	var n int

	i := 1
	j := 2
	answer := 2
	ran := 1

	fmt.Scanln(&n)

	for {
		if n == 1 {
			fmt.Println(1)
			break
		}
		ran = ran + (2 * i) + (j * 2)
		if n <= ran {
			fmt.Println(answer)
			break
		}

		i += 2
		j++
		answer++
	}
}
