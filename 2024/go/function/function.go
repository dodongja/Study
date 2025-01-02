package main

import "fmt"

//함수 return 값도 여러개 가능
//파라미터로 값만 전달(래퍼런스 전달 x)
func Multiple(a int, b int) int {
	return a * b
}

func F(n int) int {
	//2보다 작으면 - 됨
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	// 피보나치 수열 9번째 값을 출력
	fmt.Println(F(9))
}
