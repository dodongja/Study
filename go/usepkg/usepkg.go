package main

import (
	"fmt"
	"go/usepkg/custompkg" // 하위 패키지는 패키지 경로 적어 줘야함

	"github.com/guptarohit/asciigraph"         // go mod tidy를 실행해 패키지 다운로드 후 go.mod파일과 go.sum 파일에 적어주게 됨
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}
