package main

import "fmt"

var g int = 100

func main() {
	fmt.Println(g)

	{
		var s int = 50
		fmt.Println(s)
	}

	fmt.Println(s)
	
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var z int = 8
	var v int = 9
	var x int = z * v

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(x)
}
