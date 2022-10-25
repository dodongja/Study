package main

import "fmt"

/* 기본문
light := "red"

if light == "green" {			// if 조건문 {
	fmt.Println("길을 건넌다")  //    문장
} else {					    //  } else {
	fmt.Println("대기한다")     //    문장
	                           //   }
}*/

/* && ||
var age = 22

if age >= 10 && age <= 15 {
	fmt.Println("You are young")
} else if age > 30 || age < 20 {
	fmt.Println("You are not 20s")
} else {
	fmt.Println("Best age of your life")
}*/
/*쇼트서킷
var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}

	fmt.Println("cnt:", cnt)
}*/
/* 중첩 if
func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈")
		} else {
			fmt.Println("사람 없노 나눠 내자")
		}
	} else {
		fmt.Println("내가 쏜다")
	}
}*/

/*if 초기문; 조건문
func getMyAge() (int, bool) {
	return 33, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("어리노", age)
	} else if ok && age < 30 {
		fmt.Println("좋을때고", age)
	} else if ok {
		fmt.Println("아름답노", age)
	} else {
		fmt.Println("Error")
	}
}*/
// 낮 최고 25도 이상 강수확률 80% 이상 "덥고 비가 옵니다"
// 낮 최고 25도 이상 강수확률 25% 이상 "덥고 습합니다"
// 난 최고 25도 이상 강수확률 20% 미만 "야외 활동하기 좋습니다"
// 기온이 25도 이상 아니고 기온 10도 미만이거나 강수 확률이 80% 이상이면 "야외 활동하기 좋지 않습니다"
// 그렇지 "좋은 날씨 입니다"
func weather(temp, rain int) {
	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("덥고 비가 옵니다")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다")
		} else {
			fmt.Println("야외 활동하기 좋습니다")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외 활동하기 좋지 않습니다")
	} else {
		fmt.Println("좋은 날씨 입니다")
	}
}
func main() {
	weather(25, 21)

}
