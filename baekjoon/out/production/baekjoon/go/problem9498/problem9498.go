package main

import (
	"fmt"
)

func gredeCalculator(score *int) string {
	if *score >= 90 {
		return "A"
	} else if *score >= 80 {
		return "B"
	} else if *score >= 70 {
		return "C"
	} else if *score >= 60 {
		return "D"
	} else {
		return "F"
	}

}

func main() {
	var score int

	for {
		_, err := fmt.Scanln(&score)

		if err != nil {
			fmt.Println(err)
			continue
		} else if score > 100 || score < 0 {
			fmt.Println("0~100사이의 숫자를 입력해 주세요")
			continue
		} else {
			fmt.Println(gredeCalculator(&score))
			break
		}
	}
}
