package main

import (
	"fmt"
)

/*
// 가변 인수 함수
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())

}*/

/* defer
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다")

	fmt.Println("파일에 Hello World를 쏩니다")
	fmt.Fprintln(f, "Hellow World")
}
*/
/* 함수 타입 변수
func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// int타입 인수가 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperator("*")

	// 함수 타입 변수를 사용해서 함수 호출
	var result = operator(3, 4)
	fmt.Println(result)
}*/

func getOperator(op string) func(int, int) int {
	if op == "+" {
		//함수 리터럴 사용해서 더하기 함수 정의 반환
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		//함수 리터럴 사용해서 곱하기 함수 정의 반환
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)
}
