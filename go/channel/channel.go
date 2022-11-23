package main

import (
	"context"
	"fmt"
	"sync"
)

/* 채널
func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //채널 생성

	wg.Add(1)
	go square(&wg, ch) // 고루틴 생성
	ch <- 9	//채널에 9 넣음
	wg.Wait() // 작업 완료되길 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch //데이터 빼옴

	time.Sleep(time.Second) //1초 대기
	fmt.Printf("Square: %d\n", n*n)
	wg.Done() // 작업 완료
}*/

// 대기 상태
/*func main() {
	ch := make(chan int)

	ch <- 9 // mina() 함수 여기서 멈춤
	fmt.Println("COME ON")
}*/

//close()
/*func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// 데이터를 넣고 채널이 필요 없기 때문에 닫음 -> 정상종료
	close(ch)
	wg.Wait()
}*/

//select문
/*func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { //select 문
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)

		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}*/

/*func square(wg *sync.WaitGroup, ch chan int) {
	//time.Tick()은 일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환 함수
	//반환한 채널에서 데이터를 읽어오면 현재 시각을 나타내는 Time 객체 반환
	tick := time.Tick(time.Second) //1초 간격 시그널
	//현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환 함수
	//반환한 채널에서 데이터를 읽어오면 현재 시각을 나타내는 Time 객체 반환
	terminate := time.After(10 * time.Second) // 10초 이후 시그널

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}*/

/*type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go paintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			//1초마다 Body 생성
			car := &Car{}
			car.Body = "Sports car"
			//tireCh 채널에 car 넣으면 InstallTire 작동
			tireCh <- car
		case <-after:
			//10초 뒤에 Body 생성 끝
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		//1초 기다리고 Tire 생성
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		//paintCh에 car 넣으면 PaintCar 작동
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func paintCar(paintCh chan *Car) {
	for car := range paintCh {
		//1초기다리고 Color
		time.Sleep(time.Second)
		car.Color = "Red"
		//현재시각에서 시작시각을 빼서 경과시간 출력
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(),
			car.Body, car.Tire, car.Color)
	}
	wg.Done()
}*/

/*
func main() {
	wg.Add(1)
	// ctx 컨텍스트 객체
	//cancel 취소 함수를 사용해서 원할 때 취소할 수 있음
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go PrintEverySecond(ctx)
	//5초 뒤에 cancel()
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		//인수로 받은 컨텍스트의 Done() 채널의 시그널이 있는지를 검사
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}*/

var wg sync.WaitGroup

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)

	go square(ctx)

	wg.Wait()
}
