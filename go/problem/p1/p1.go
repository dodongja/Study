package main

import "fmt"

func main() {
	var name string
	var age int
	var weight float32

	n, err := fmt.Scanln(&name, &age, &weight)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Printf("내이름은 %s, 나이는 %d, 몸무게는 %.2f", name, age, weight)
	}

}
