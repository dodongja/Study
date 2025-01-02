package main

import "fmt"

func main() {
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)

	var a int8 = 127
	var b int8 = a + 1
	//오버플로(overflow)
	//127 01111111
	//127 + 1 = 0111111 + 1
	//-128 = 1000000
	fmt.Println(a)
	fmt.Println(b)

	var c int8 = -127
	var d int8 = c - 1

	//언더플로(underflow)
	//-128 1000000
	//-128 - 1= 1000000 - 0000001
	//127 = 01111111
	fmt.Println(c)
	fmt.Println(d)

	var e float64 = 0.1
	var f float64 = 0.2
	var g float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", e, f, g, e+f == g)
	fmt.Println(e + f)

	var h int8 = 30
	//복합 대입연산자
	// h = h << 2
	h <<= 2
	// h = h + 8
	h += 8

}
