package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var b int
	var arr []int
	var answer string

	fmt.Scanln(&n, &b)

	if b > n {
		if n >= 10 {
			answer = string(n + 55)
		} else {
			answer = strconv.Itoa(n)
		}
	} else {
		for {
			if b > n {
				break
			}

			remainder := n % b
			arr = append(arr, remainder)
			n = n / b

			if b > n {
				arr = append(arr, n)
			}

		}

		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i] >= 10 {
				answer += string(arr[i] + 55)
			} else {
				answer += strconv.Itoa(arr[i])
			}
		}
	}

	fmt.Print(answer)

}
