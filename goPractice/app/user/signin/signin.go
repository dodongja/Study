package signin

import (
	"encoding/json"
	"errors"
	"fmt"
	"goPractice/app/user/jwt"
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
	var user signup.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dsn := "eddi:eddi@123@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var compareUser signup.User
	//first / find 뭔차이고...
	findUser := db.First(&compareUser, "email = ?", user.Email)
	fmt.Println(user.Email)
	if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
		resp := make(map[string]string)
		resp["message"] = "해당 ID가 존재하지 않습니다"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	if signup.VerifyPassword(compareUser.Password, user.Password) != nil {
		resp := make(map[string]string)
		resp["message"] = "Password를 확인하세요"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}
	//뭔가 코드가 반복됨 흠..
	if _, err := jwt.CreateAccessToken(w, &compareUser); err != nil {
		resp := make(map[string]string)
		resp["message"] = "token can't create"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	if _, err := jwt.CreateRefreshToken(w, &compareUser); err != nil {
		resp := make(map[string]string)
		resp["message"] = "token can't create"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	//response body에 json 데이터 보내는거 다른 방법 찾아봐야할듯
	resp := make(map[string]string)
	resp["message"] = "success"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)

}
