package main

import (
	"fmt"
)

func main() {
	/*
		var x int32 = 7
		var y int32 = 3

		var s float32 = 3.14
		var t float32 = 5

		fmt.Println("x + y =", x+y)
		fmt.Println("x - y =", x-y)
		fmt.Println("x * y =", x*y)
		fmt.Println("x / y =", x/y)
		fmt.Println("x % y =", x%y)

		fmt.Println("s * t =", s*t)
		fmt.Println("s / t =", s/t)*/

	/*비트 연산자
	var x1 int8 = 34
	var x2 int16 = 34
	var x3 uint8 = 34
	var x4 uint16 = 34

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, uint8(^x3))
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, uint16(^x4))*/

	/*왼쪽 시프트
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)*/

	/*오른쪽 시프트
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)
	fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", uint8(y), uint8(y>>2), y>>2)
	fmt.Printf("z:%08b x>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w:%08b x>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), x>>2)*/

	/*오버 플로우 언더 플로우
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t = %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t = %4d, %08b\n", x+1, x+1)
	fmt.Printf("x + 2\t = %4d, %08b\n", x+2, x+2)
	fmt.Printf("x + 3\t = %4d, %08b\n", x+3, x+3)

	var y int8 = -128
	fmt.Printf("%d > %d - 1 : %v\n", y, y, y > y-1)
	fmt.Printf("y\t = %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t = %4d, %08b\n", y-1, y-1)*/

	/* 실수 오차
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)*/

	/* math.Nextafter
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	//
	fmt.Println(math.Nextafter(a, b))
	fmt.Println(math.Nextafter(b, a))
	fmt.Println(math.Nextafter(a, c))
	fmt.Println(math.Nextafter(a+b, c))
	fmt.Println(math.Nextafter(a+b, c) == c)
	*/

	/*math/big
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // c와 d를 비교. c가 작을경우 -1 c가 큰 경우 1 , 같을경우 0*/

	var a int = 10
	var b int = 20


}
