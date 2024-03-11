package main

import "fmt"

func main() {
	var n string
	var b int
	var answer int

	number := map[string]int{
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
		"G": 16,
		"H": 17,
		"I": 18,
		"J": 19,
		"K": 20,
		"L": 21,
		"M": 22,
		"N": 23,
		"O": 24,
		"P": 25,
		"Q": 26,
		"R": 27,
		"S": 28,
		"T": 29,
		"U": 30,
		"V": 31,
		"W": 32,
		"X": 33,
		"Y": 34,
		"Z": 35,
	}

	fmt.Scanln(&n, &b)

	numberLength := len(n)

	for i := 0; i < numberLength; i++ {
		placeValue := 1
		if number[string(n[i])] == 0 {
			for j := 0; j < numberLength-1-i; j++ {
				placeValue *= b
			}
			placeValue *= int(n[i])
		} else {
			for j := 0; j < numberLength-1-i; j++ {
				placeValue *= b
			}
			placeValue *= number[string(n[i])]
		}
		answer += placeValue
	}

	fmt.Print(answer)
}
