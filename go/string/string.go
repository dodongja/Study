package main

import "fmt"

func main() {
	/*poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워 했다.\n"

		poet2 := `죽는 날까지 하늘을 우러러
	한 점 부끄럼이 없기를,
	잎새에 이는 바람에도
	나는 괴로워 했다`

		fmt.Println(poet1)
		fmt.Println(poet2)*/

	/*var char rune = '한'

	fmt.Printf("%T\n", char) // 타입 int 32
	fmt.Println(char)  // int32라 숫자로 나옴 54620
	fmt.Printf("%c\n", char) // 한*/

	//str := "Hello 월드!"
	//arr := []rune(str)
	/*for i := 0; i < len(arr); i++ {
		fmt.Printf("%c\n", arr[i])
	}*/

	/*for _, v := range str {
		fmt.Printf("%c\n", v)
	}*/

	str1 := "안녕하세요"
	str2 := str1
	fmt.Println(str2)
	str1 = "안녕히가세요"
	fmt.Println(str1)
	fmt.Println(str2)

}
