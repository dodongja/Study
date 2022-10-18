package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"os" // 표준 입출력 등을 가지고 있는 패키지
)

func main() {
	stdin := bufio.NewReader(os.Stdin) //표준 입력을 읽는 객체
	var i int
	var k int // 값을 저장할 변수

	n, err := fmt.Scanln(&i, &k)
	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n') // 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, i, k)
	}
	n, err = fmt.Scanln(&i, &k)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, i, k)
	}
}
