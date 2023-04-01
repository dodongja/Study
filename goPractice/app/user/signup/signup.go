package signup

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	//"github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	//password 인코딩 패키지 필요할 듯
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

// 회원 가입
// 핸들러 부분이랑 로직 부분 나누는듯?
// 아이디 / 닉네임 중복 방지 해야함
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Password, _ = HashPassword(user.Password)

	dsn := "eddi:eddi@123@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//	gorm은 구조체 이름의 마지막에 "s"를 붙여서 테이블을 만든다.
	// created_at, updated_at, deleted_at 필드를 기본으로 생성한다.
	// auto_increment 옵션을 가진 id 필드를 기본으로 생성한다.
	//테이블 생성 이래 계속 나둬도 되나?
	db.AutoMigrate(&User{})

	//unique 에러가 뜨는데 그게 email, nickname인지 어케 아노??
	//Error 1062 (23000): Duplicate entry 'dongja' for key 'users.nickname'
	//Error 1062 (23000): Duplicate entry 'qs553@nate.com' for key 'users.email'
	//뭔 칼럼이 에러다라고 메서드해서 못가져오나...?
	//그럼 결국 unique 먹이는게 더 좋을라나
	// 앞단에 db에 값 가져와서 중복체크부터하고 Create 해야할듯??
	//errors.Is(db.Create(&user).Error, ErrDuplicatedKey)
	//unique 포기 걍 앞단에서 중복체크 이게맞나...
	if checkEamilDuplication(db, &user) {
		response(w, "email 중복")
	} else if checkNicknameDuplication(db, &user) {
		response(w, "nickname 중복")
	} else {
		db.Create(&user)
		response(w, "가입 성공")
	}

}

// func DuplicateEmail()(){

// }
// 에러 핸들링 어렵네... 이게 맞나..
// db에 값 가져와서 변수에 어케 박노
func checkEamilDuplication(db *gorm.DB, user *User) bool {
	check := db.First(&user, "email = ?", user.Email)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		return false
	} else if check != nil {
		return true
	}
	return false
}

func checkNicknameDuplication(db *gorm.DB, user *User) bool {
	check := db.First(&user, "nickname = ?", user.Nickname)
	if errors.Is(check.Error, gorm.ErrRecordNotFound) {
		return false
	} else if check != nil {
		return true
	}
	return false
}

func response(w http.ResponseWriter, msg string) {
	resp := make(map[string]string)
	resp["message"] = msg
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
