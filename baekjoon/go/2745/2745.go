package main

import "fmt"

func main() {
	var n string
	var b int
	var answer int

	number := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
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

	fmt.Println("n[iii]", string(n[0]))

	for i := 0; i < numberLength; i++ {
		placeValue := 1

		for j := 0; j < numberLength-1-i; j++ {
			placeValue *= b
		}

		answer += (placeValue * number[string(n[i])])
	}

	fmt.Print(answer)
}
