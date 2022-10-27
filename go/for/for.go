package main

import "fmt"

func main() {
	/* 기본 for문
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	*/
	// fmt.Println(i)   error

	/* 무한루프
	i := 1
	for {
		time.Sleep(time.Second * 2) // 2초 지연
		fmt.Println(i)
		i++
		// ctrl + c 강제 종료
	}*/

	/* continue break
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요")

			stdin.ReadString('\n') //키보드 버퍼를 비움
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break
		}
	}

	fmt.Print("for문이 종료 됩니다.")*/
	// 1 2 3 4 5 4 3 2 1
	// 1 2 3 4 5 6 7 8 9
	// 2 3 4 5 6 7 8 9 10
	// k 1 2 3 4 5
	// k 4 3 2 1
	//  k -1

	/* 중첩 for문
	var k int = 0
	const length int = 10
	for i := 0; i < length; i++ {
		if i >= (length / 2) {
			k -= 1
		} else {
			k += 1
		}
		for j := 0; j < k; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}*/

/*	a := 1
	b := 1

OuterFor: //레이블 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor 
			}
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
*/


}
