package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { //매개 변수로 Data 포인터로 받음
	arg.value = 999
	arg.data[100] = 999

	fmt.Printf("arg %p\n", arg)
}

func main() {
	var data Data

	ChangeData(&data) // 인수로 data 주소를 넘김
	var p2 *Data = &data

	fmt.Printf("%p\n", p2)
}
