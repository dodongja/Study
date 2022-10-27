package main

import "fmt"

func main() {
	/*var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // var 변수명 [요소 개수]타입

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}*/
	/* range
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.20}

	for i, v := range t { // i = index , v = 원솟값
		fmt.Print(i, v)
	}

	for _, v := range t { // '_' 이용해서 인덱스를 무효화 할 수 있음
		fmt.Print(v)
	}*/
	/* 배열 복사
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{100, 200, 300, 400, 500}

	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	b = a // a배열을 b배열에 복사

	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}*/

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 여러 줄에 걸쳐 초기화 할때 , 찍음. 추후 항목이 늘어날 경우 쉼표를 찍지 않아서 생길 수 있는 오류를 방지
	}

	b := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}} // 마지막에 닫히는 중괄호가 마지막 항목과 같은줄에 있으면 쉼표를 찍지 않아도 됨

	for _, arr := range a {
		for _, v := range arr {
			fmt.Println(v)
		}
	}

	for _, arr := range b {
		for _, v := range arr {
			fmt.Println(v)
		}
	}

}
