package main

import (
	"encoding/json"
	"net/http"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	email    string `gorm:"column:email"`
	password string `gorm:"column:password"`
	//password 인코딩 패키지 필요할 듯
}

func MakeHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/signup", SignUpHandler).Methods("GET")

	return mux
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user User = User{
		email:    "qw553@nate.com",
		password: "qw553",
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dsn := "eddi:eddi@123@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//	gorm은 구조체 이름의 마지막에 "s"를 붙여서 테이블을 만든다.
	// created_at, updated_at, deleted_at 필드를 기본으로 생성한다.
	// auto_increment 옵션을 가진 id 필드를 기본으로 생성한다.
	//db.CreateTable(&User{})
	db.Create(&user)
}

func main() {
	http.ListenAndServe(":3000", MakeHandler())

}
