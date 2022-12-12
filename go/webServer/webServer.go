package main

import (
	"fmt"
	"net/http"
	"strconv"
)
/*
func main() {
	//포트 3000번에 "/"경로에 수신 시 호출하는 핸들러
	//두번째 인수 실행할 함수
	//이 때 함수는 반드시 http.ResponseWriter를 받고 두 번째 인수로 *http.Request 타입을 받아야 함

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Print() 함수가 표준 출력 스트림으로 출력이 고정
		//Fprint() 지정한 출력 스트림에 출력한다는 점이 다름
		//여기서는 http.ResponseWriter 타입을 출력 스트림으로 지정
		//http.ResponseWriter 타입에 값을 쓰면 HTTP 응답으로 전송
		fmt.Fprint(w, "Hello World")
	})

	//포트 번호 -> 요청을 기다릴 주소
	//nil -> DefaultServeMux 사용 http.HandleFunc()함수를 호출 해 웹 서버를 실행

	http.ListenAndServe(":3000", nil)
}*/

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == ""{
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name , id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}


