package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	os.Remove(filename)
	line, err := ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("파일 생성 실패", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기 실패", err)
			return
		}
	}

	fmt.Println("파일내용:", line)
}*/
/*
type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw")
	fmt.Println(err)
	if err != nil {
		//err 타입을 PasswordError로 변환시켜서 에러메시지 + Len과 RequireLen도 접근 가능
		// err.(type) 인듯 type assertion?
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}*/

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	return a * b, nil

}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	fmt.Println(scanner.Text())
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
