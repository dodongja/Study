package main

import (
	"fmt"
	"math"

	// "math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	// //Nextafter() : float64 타입 두개를 받아서 float64 타입 하나 반환, x에서 y를 향해서 1비트 조정한 값 반환
	// var a float64 = 0.1
	// var b float64 = 0.2
	// var c float64 = 0.3

	// fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	// fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	// a = 0.0000000000004
	// b = 0.0000000000002
	// c = 0.0000000000007
	// //어디까지나 오차가 매우 작을뿐이지, 정확한 계산은 아님
	// fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

	//금융 프로그램이면 math/big 패키지에서 제공하는 float객체를 사용해야함
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	//0이면 두 값이 같을 경우
	fmt.Println(c.Cmp(d))
}
