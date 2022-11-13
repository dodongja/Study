package main

import (
	"time"
)

/*type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	//리시버가 포인터 변수고 그냥 일반 변수이나 mainA.withdrawValue 같이 사용 가능 한 이유는 Go언어에서 자동으로 변환 시켜주기 때문이다.
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 리시버가 포인터기 때문에 a1이 mainA주소를 가리키기 때문에 balance = 70

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 리시버가 포인터기 때문에 a1이 mainA주소를 가리키기 때문에 balance = 50

	var mainB account = mainA.withdrawReturnValue(20) // mainB에 mainA의 값들이 복사됨 (주소 값 x)
	fmt.Println(mainB.balance) // 50

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance) // 20
}*/

//연습문제

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt         *Product
	ShippedTime time.Time
	DeliverTime time.Time
}

func (c *Courier) sendProduct(p *Product) *Parcel {
	par := &Parcel{}
	par.Pdt = p
	par.ShippedTime = time.Now()

	return par
}

func (p *Parcel) Delivered() *Product {
	p.DeliverTime = time.Now()
	return p.Pdt
}
