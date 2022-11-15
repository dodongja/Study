package main

import (
	"sort"
)

/*func changeSlice(slice2 []int) {
	slice2[2] = 200
}*/

//func main() {
/*//slice사용 주의점
//slice := make([]int, 3, 5)
slice := []int{1, 2, 3}

slice2 := append(slice, 4, 5) //이건 배열의 크기보다 큰 값이 들어가서 새로운 배열을 만들어서 추가한것이라 그럼.

fmt.Println(slice)
fmt.Println(slice2)

fmt.Printf("%p\n", &slice[0])
fmt.Printf("%p\n", &slice2[0])

slice[1] = 5

fmt.Println(slice)
fmt.Println(slice2)*/

//배열 슬라이싱
/*array := [5]int{1, 2, 3, 4, 5}

slice := array[1:2]

fmt.Println(cap(slice))
fmt.Printf("%p\n", &slice[0])
fmt.Printf("%p\n", &array[1])

slice = append(slice, 500)

fmt.Println(slice)
fmt.Println(array)

slice = append(slice, 400, 300, 200)

fmt.Println(slice)
fmt.Println(array)

fmt.Printf("%p\n", &slice[0])
fmt.Printf("%p\n", &array[1])*/

//슬라이스 슬라이싱
/*slice1 := []int{1, 2, 3}
slice2 := slice1[:]

fmt.Println(slice2)*/

//슬라이스 복제 같은 배열 안가리키게 하기
//slice1 := []int{1, 2, 3, 4, 5}

//append() 함수로 코드 개선
/*slice2 := append([]int{}, slice1...)

slice2[0] = 500

fmt.Println(slice1)
fmt.Println(slice2)*/

/*s := []int{5, 2, 6, 3, 1, 4}
sort.Ints(s)

fmt.Println(s)*/

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18},
	}

	sort.Sort(Students(s))
	
}
