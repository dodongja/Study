package main

import "fmt"

func main() {
	var n int
	var answer int
	var multiple int
	var firstNum int
	const num = 2

	firstNum = 2

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		if i == 0 {
			answer += firstNum
			multiple = 1
		}

		answer += multiple
		multiple *= num
	}

	fmt.Println(answer * answer)
}
