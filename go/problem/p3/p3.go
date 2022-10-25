package main

import "fmt"

func main() {
	var x string
	var y string

	_, err := fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s의 문자열의 길이는 %d이고 %s의 문자열의 길이는 %d 입니다", x, len(x), y, len(y))
	}

}
