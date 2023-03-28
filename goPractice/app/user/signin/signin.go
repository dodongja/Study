package signin

import (
	"encoding/json"
	"fmt"
	"goPractice/app/user/signup"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//로그인 함 -> 아이디 비밀번호
//db에서 아이디 비밀번호 조회 해야함
//비밀번호는 디코딩 해서 비교해야함
//비밀번호 틀린 경우 / 아이디 존재하지 않는 경우
//나중에는 oauth email도 db저장하면서 회원가입할 때 oauth/signup으로 나눠야할듯

type SignInUser struct {
	gorm.Model
	Email    string `json:"email"  gorm:"unique"`
	Password string `json:"password"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInUser SignInUser

	err := json.NewDecoder(r.Body).Decode(&signInUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	checkUser(&signInUser)

}

func checkUser(signInUser *SignInUser) {
	var user signup.User
	dsn := "eddi:eddi@123@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Print(db.First(&user, "Email = ?", signInUser.Email))
}
