package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numCheck() {
	// 0 <= n < 100
	rand.Seed(time.Now().UnixNano())

	var answer int = rand.Intn(100)
	fmt.Println("answer ", answer)
	var input int
	var count int

	for {
		fmt.Print("숫자 값을 입력하세요 : ")
		fmt.Scanln(&input)
		count++

		if input == answer {
			fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수 : %d", count)
			break
		} else if input > answer {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		}
	}
}
func main() {
	//맞춰야 하는 숫자 정하기
	//숫자 입력
	//숫자 맞는지 확인하기
	//예외처리가 필요함
	//time써야 rand 숫자가 다시 run했을시 변경 됨 아니면 계속 유지
	numCheck()

}
