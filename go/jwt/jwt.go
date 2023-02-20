package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	//password 인코딩 패키지 필요할 듯
}

func main() {
	//파일 구조를 어케하노
	//일단 http request -> db connect -> httpresponse
	//http 통신 하는 부분이랑 db 연동하는 부분 코드 나누면 될듯??
	/*db, err := sql.Open("mysql", "eddi:eddi@123@tcp(localhost:3306)/golang_db")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	var name string
	err = db.QueryRow("SELECT email FROM user WHERE id = 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
	print("success")*/
	/*var user User = User{
		email:    "qw553@nate.com",
		password: "qw553",
	}*/

	dsn := "eddi:eddi@123@tcp(localhost:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//	gorm은 구조체 이름의 마지막에 "s"를 붙여서 테이블을 만든다.
	// created_at, updated_at, deleted_at 필드를 기본으로 생성한다.
	// auto_increment 옵션을 가진 id 필드를 기본으로 생성한다.
	//table을 그냥 mysql에 직접 만드는건가??
	//fmt.Print(user)
	db.AutoMigrate(&User{})
	//db.Create(&user)
}
