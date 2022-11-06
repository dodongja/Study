package main

import (
	"fmt"
)

/*
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "서울시 강동구 ..."
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소: ", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입: ", house.Type)
}*/
/*
type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}*/
/*
type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23, 2}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 2},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 가격: %d만 원 유저 레벨: %d\n",
		vip.Name, //하나로 접근할 수 있다
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
		vip.User.Level, // 필드명이 중복일 때 포함된 구조체명을 쓰고 접근
	)
}*/

type User struct {
	Age   int32
	Score float64
}

func main() {
	var user = User{15, 23.3}

	user2 := user

	var p1 *User = &user
	var p2 *User = &user2

	fmt.Printf("%p\n", p1)
	fmt.Printf("%p\n", p2)
}
