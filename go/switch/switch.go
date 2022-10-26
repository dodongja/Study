package main

import "fmt"

/* 여러 값 비교
day := "thursday"

switch day {
case "monday", "tuesday": //여러 값 검사 가능
	fmt.Println("월, 화요일은 수업 가는 날입니다.")
case "wednesday", "thursday", "friday":
	fmt.Println("수,목,금요일은 실습 가는 날입니다")
}*/

/*조건문 비교
temp := 18

switch true {	// 비굣값과 case의 값이 같아지는 경우를 찾는 구문이기 때문에 비굣값을 true로 함
case temp < 10, temp > 30:
	fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다")
case temp >= 10 && temp < 20:
	fmt.Println("약간 추울 수 있으니 가벼운 겉옷 준비")
case temp >= 15 && temp < 25:
	fmt.Println("야외활동하기 굳")
default:
	fmt.Println("따뜻")
}*/
func getMyAge() int {
	return 22
}

//func main() {
/*switch age := getMyAge(); true {
case age < 10:
	fmt.Println("child")
case age < 20:
	fmt.Println("Teenager")
case age < 30:
	fmt.Println("20s")
default:
	fmt.Println("My age is", age)
}*/
/*switch age := getMyAge(); age {
case 10:
	fmt.Println("Teenage")
case 33:
	fmt.Println("Pair 3")
default:
	fmt.Println("My age", age)
}*/

//fmt.Println("age is", age) error - age 변수는 사라졌습니다
//}

/*type ColorType int
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}*/

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}
func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315:
		return North
	case angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle > 135 && angle < 225:
		return South
	case angle >= 255 && angle < 315:
		return West
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
}
