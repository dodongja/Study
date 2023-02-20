package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rs/cors"

	//"github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	//password 인코딩 패키지 필요할 듯
}

func MakeHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/signup", SignUpHandler).Methods("POST")

	return mux
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user User

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
	//테이블 생성 이래 계속 나둬도 되나?
	db.AutoMigrate(&User{})
	fmt.Println("************************************")
	db.Create(&user)
	message := "가입 성공"
	w.Write([]byte(`{"message": "` + message + `"}`))

}

func main() {
	//cors도 따로 뺴야될듯?
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	http.ListenAndServe(":3000", cors.Handler(MakeHandler()))

}
