package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var standaradPieces [6]int = [6]int{1, 1, 2, 2, 2, 8}
	var piece int
	var slice []int

	for i := 0; i < len(standaradPieces); i++ {
		_, err := fmt.Scan(&piece)
		if err != nil {
			fmt.Println(err)
			stdin.ReadString('\n')
		} else {
			slice = append(slice, piece)
		}
	}

	for i := 0; i < len(standaradPieces); i++ {
		fmt.Print(standaradPieces[i]-slice[i], " ")
	}

}
