package main

import "fmt"

func main() {
	var x int
	var y int

	n, err := fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Printf("나머지 값은 %d 입니다", y%x)
	}

}
