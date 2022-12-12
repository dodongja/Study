package main

import "fmt"

func main() {
	var year int

	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Println(err)
	} else {
		if year%4 == 0 {
			if year%400 == 0 {
				fmt.Println(1)
			} else if year%100 == 0 {
				fmt.Println(0)
			} else {
				fmt.Println(1)
			}
		} else {
			fmt.Println(0)
		}
	}
}
