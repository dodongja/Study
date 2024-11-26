package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297
	//기본 서식에 맞춰 표준 출력
	fmt.Print("a: ", a, "b: ", b)
	//기본 서식에 맞춰 표준 출력 후 개행
	//실숫값 기본 서식 %g
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	//주어진 사용자 서식에 맞춰 입력값 출력
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f)

	str := "Hello\tGO\t\tWorld\n\"Go\"is Awesome!\n"
	//\t tab \n 개행 \" " \\ \
	fmt.Print(str)

	var c int
	var d int
	// & 변수의 메모리 주소
	// Scanln은 반드시 입력 받을 때 enter키로 입력을 종료해야함
	n, err := fmt.Scanln(&c, &d)
	//err nill 아니면 에러
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, c, d)
	}
}
